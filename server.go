package main

import "github.com/paulmj7/hermes/hermes"

func main() {
	port := ":3000"
	roots := []string{"/"}
	hiddenFiles := make(map[string]bool)
	corsEnabled := false
	w := hermes.Worker{port, roots, hiddenFiles, corsEnabled}
	w.Serve()
}
