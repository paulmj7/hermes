package hermes

import (
	"log"
	"net/http"
)

type Worker struct {
	Port   string
	Roots  []string
	Hidden map[string]bool
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

func (w Worker) Serve() {
	http.HandleFunc("/api", w.index)
	http.HandleFunc("/api/change_dir", w.changeDir)
	http.HandleFunc("/api/retrieve", getFile)
	http.HandleFunc("/api/send", sendFile)
	http.HandleFunc("api/upload", saveFile)
	http.HandleFunc("api/create", createFolder)
	http.HandleFunc("/api/move", move)
	http.HandleFunc("/api/delete", delete)
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
