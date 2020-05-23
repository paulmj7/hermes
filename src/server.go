package main

import (
	"fmt"
	"log"
	"net/http"
)

var roots []string

func main() {
	config := ReadConfig("config.json")
	roots = ToString(config.Roots)
	http.HandleFunc("/api", Index)
	http.HandleFunc("/api/move", ChangeDir)
	http.HandleFunc("/api/retrieve", GetFile)
	http.HandleFunc("/api/send", SendFile)
	fmt.Println("Listening on " + config.Port)
	log.Fatal(http.ListenAndServe(config.Port, nil))
}

func SetupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-ZMethods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
