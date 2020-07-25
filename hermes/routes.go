package hermes

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func (worker *Worker) index(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	rootbrowser := []Item{}
	i := 1

	for _, item := range worker.Roots {
		name, _ := SplitPath(item, "/")
		tempItem := Item{Name: name, Path: item, Parent: "", Root: item, IsFile: false, Size: -1, DateMod: "", ID: i}
		rootbrowser = append(rootbrowser, tempItem)
		i++
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rootbrowser)
}

func (worker *Worker) changeDir(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	}

	decoder := json.NewDecoder(r.Body)
	var req ReqBody
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !FromRoot(req.Path, req.Root) {
		fmt.Println(req.Path + " not available from root: " + req.Root)
		return
	}

	files, er := ioutil.ReadDir(req.Path)
	if er != nil {
		fmt.Println(er)
		return
	}

	contents := []Item{}
	i := 1

	for _, f := range files {
		destination := path.Join(req.Path, f.Name())
		if worker.Hidden[destination] {
			continue
		}
		isFile, size, dateMod, e := ItemInfo(destination)
		if e != nil {
			fmt.Println(e)
			return
		}

		tempItem := Item{Name: f.Name(), Path: destination, Parent: req.Path, Root: req.Root, IsFile: isFile, Size: size, DateMod: dateMod, ID: i}
		contents = append(contents, tempItem)
		i++
	}

	fmt.Println(r.RemoteAddr + " accessed " + req.Path)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contents)
}

func getFile(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req ReqBody
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := []byte(req.Path)
	str := base64.StdEncoding.EncodeToString(data)
	resBody := ResBody{Key: str}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resBody)
}

func sendFile(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		fmt.Println("Request missing parameter")
		return
	}

	key := keys[0]
	data := string(key)
	str, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	path := string(str)
	name, _ := SplitPath(path, "/")
	fmt.Println("Sending file: " + name)
	w.Header().Set("Content-Disposition", "attachment; filename="+name)

	f, er := os.Open(path)
	if er != nil {
		return
	}

	defer f.Close()
	fi, e := f.Stat()
	if e != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	http.ServeContent(w, r, path, fi.ModTime(), f)
}

func saveFile(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(32 << 20)
	destination := r.FormValue("path")
	fi, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fi.Close()

	f, e := os.Create(path.Join(destination + "/" + handler.Filename))
	if e != nil {
		fmt.Println(e)
		return
	}
	wt := bufio.NewWriter(f)

	defer f.Close()
	n, er := io.Copy(wt, fi)
	fmt.Println("write", n)
	if er != nil {
		fmt.Println(er)
		return
	}
	wt.Flush()
	w.WriteHeader(http.StatusCreated)
}

func createFolder(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req CreateReqBody
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.MkdirAll(req.Path, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Create folder at path: " + req.Path)
	w.WriteHeader(http.StatusCreated)
}

func move(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	if r.Method != http.MethodPut {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req map[string]string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	fmt.Println("Move " + req["location"] + " to " + req["destination"])
	err = os.Rename(req["location"], req["destination"])
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func delete(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	if r.Method != http.MethodDelete {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req map[string]string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.RemoveAll(req["path"])
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
