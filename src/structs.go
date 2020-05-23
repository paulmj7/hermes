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
	Port  string `json:"port"`
	Roots []Root `json:"roots"`
}

type Root struct {
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

func ToString(roots []Root) []string {
	var extracted []string
	for _, item := range roots {
		extracted = append(extracted, item.Path)
	}

	return extracted
}
