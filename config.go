package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Paths struct {
		Projects string `json:"projects"`
		SSL      string `json:"ssl"`
		Vhost    string `json:"vhost"`
	} `json:"paths"`
	Template struct {
		NoSSL string `json:"nossl"`
		SSL   string `json:"ssl"`
	} `json:"template"`
	Dirs struct {
		Ignore []string `json:"ignore"`
	} `json:"dirs"`
}

func readFile(cfg *Config) {
	f, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	jsonParser := json.NewDecoder(f)
	jsonParser.Decode(&cfg)
}
