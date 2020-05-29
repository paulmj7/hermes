package main

import (
	"encoding/json"
	"log"
	"os"
)

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
	Path      string `json:"path"`
	Root      string `json:"root"`
	Direction int    `json:"direction"`
}

type ResBody struct {
	Key string `json:"key"`
}

type Config struct {
	Port      string   `json:"port"`
	Roots     []Root   `json:"roots"`
	Blacklist []BLPath `json:"blacklist"`
}

type Root struct {
	Path string `json:"path"`
}

type BLPath struct {
	Path string `json:"path"`
}

func ReadConfig(configPath string) Config {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal("Error reading config")
	}

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error parsing json")
	}

	return config
}

func RootsStrings(roots []Root) []string {
	var extracted []string
	for _, item := range roots {
		extracted = append(extracted, item.Path)
	}

	return extracted
}

func BLStrings(bl []BLPath) []string {
	var extracted []string
	for _, item := range bl {
		extracted = append(extracted, item.Path)
	}

	return extracted
}

func BLMap(bl []BLPath) map[string]bool {
	m := make(map[string]bool)
	for _, item := range bl {
		m[item.Path] = true
	}

	return m
}
