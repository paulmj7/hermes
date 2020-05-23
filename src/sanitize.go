package main

import (
	"strings"
)

func FromRoot(path, root string) bool {
	l := len(root)

	if l > len(path) {
		return false
	}

	rootOfPath := path[0:l]

	if rootOfPath == root {
		return true
	}

	return false
}

func SplitPath(path, sep string) (string, string) {
	idx := strings.LastIndex(path, sep)

	curr := path[idx+1:]
	parent := path[:idx]

	return curr, parent
}
