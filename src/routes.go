package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	rootbrowser := []Item{}
	i := 1

	for _, item := range roots {
		name, _ := SplitPath(item, "/")
		tempItem := Item{Name: name, Path: item, Parent: "..", Root: item, IsFile: false, Size: -1, DateMod: "", ID: i}
		rootbrowser = append(rootbrowser, tempItem)
		i++
	}

	json.NewEncoder(w).Encode(rootbrowser)
}

func ChangeDir(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
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
	}

	files, er := ioutil.ReadDir(req.Path)
	if er != nil {
		fmt.Println(er)
	}

	contents := []Item{}
	i := 1

	for _, f := range files {
		path := req.Path + "/" + f.Name()
		isFile, size, dateMod, e := ItemInfo(path)
		if e != nil {
			fmt.Println(e)
			return
		}

		tempItem := Item{Name: f.Name(), Path: path, Parent: req.Path, Root: req.Root, IsFile: isFile, Size: size, DateMod: dateMod, ID: i}
		contents = append(contents, tempItem)
		i++
	}

	fmt.Println(r.RemoteAddr + " accessed " + req.Path)
	json.NewEncoder(w).Encode(contents)
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	if (*r).Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req ReqBody
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
	}

	data := []byte(req.Path)
	str := base64.StdEncoding.EncodeToString(data)
	resBody := ResBody{Key: str}
	json.NewEncoder(w).Encode(resBody)
}

func SendFile(w http.ResponseWriter, r *http.Request) {
	SetupResponse(&w, r)

	if (*r).Method == "OPTIONS" {
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

	http.ServeContent(w, r, path, fi.ModTime(), f)
}
