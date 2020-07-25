package hermes

import (
	"log"
	"net/http"
)

type Worker struct {
	Port   string
	Roots  []string
	Hidden map[string]bool
	CORS   bool
}

func (w Worker) WatchRoot(path string) {
	if _, ok := w.Hidden[path]; ok {
		return
	}

	w.Roots = append(w.Roots, path)
}

func (w Worker) HideRoot(path string) {
	if ok, idx := contains(w.Roots, path); ok {
		w.Roots[idx] = w.Roots[len(w.Roots)-1]
		w.Roots[len(w.Roots)-1] = ""
		w.Roots = w.Roots[:len(w.Roots)-1]
	}

	w.Hidden[path] = true
}

func (w Worker) Listen(route string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(route, handler)
}

func (w Worker) Serve() {
	http.HandleFunc("/api", w.index)
	http.HandleFunc("/api/change_dir", w.changeDir)
	http.HandleFunc("/api/retrieve", w.getFile)
	http.HandleFunc("/api/send", w.sendFile)
	http.HandleFunc("api/upload", w.saveFile)
	http.HandleFunc("api/create", w.createFolder)
	http.HandleFunc("/api/move", w.move)
	http.HandleFunc("/api/delete", w.delete)
	log.Fatal(http.ListenAndServe(w.Port, nil))
}

func contains(s []string, p string) (bool, int) {
	for k, v := range s {
		if p == v {
			return true, k
		}
	}

	return false, -1
}

type Item struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Parent  string `json:"parent"`
	Root    string `json:"root"`
	IsFile  bool   `json:"isfile"`
	Size    int64  `json:"size"`
	DateMod string `json:"datemod"`
	ID      int    `json:"id"`
}

type ReqBody struct {
	Path string `json:"path"`
	Root string `json:"root"`
}

type CreateReqBody struct {
	Path string `json:"path"`
}

type ResBody struct {
	Key string `json:"key"`
}

type Root struct {
	Path string `json:"path"`
}
