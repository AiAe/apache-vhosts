package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var Cfg Config

var configPath = userPath() + "/apache-vhosts"

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

func ReadFile(cfg *Config) {
	f, err := os.Open(configPath + "/config.json")
	if err != nil {
		createFile()
		ReadFile(cfg)
	}
	defer f.Close()

	jsonParser := json.NewDecoder(f)
	jsonParser.Decode(&cfg)
}

func createFile() {
	fmt.Println("Config file not found! Creating...")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		os.MkdirAll(configPath, 0700) // Create your file
	}

	_, err := os.OpenFile(configPath+"/config.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	createConfig()
}

func createConfig() {
	config := Config{}

	config.Template.NoSSL = "<VirtualHost *:80>\n  DocumentRoot \"%v\"\n  ServerName %v.test\n  </VirtualHost>\n"
	config.Template.SSL = "<VirtualHost *:443>\n  DocumentRoot \"%[1]v\"\n  ServerName %[2]v.test\n  SSLEngine on\n  SSLCertificateFile \"%[3]v/%[2]v.pem\"\n  SSLCertificateKeyFile \"%[3]v/%[2]v-key.pem\"\n  Protocols h2 http/1.1\n  </VirtualHost>\n"
	config.Dirs.Ignore = []string{"projects", "ssl"}

	fmt.Println("Enter projects path: ")
	fmt.Scanln(&config.Paths.Projects)
	fmt.Println("Enter ssl path: ")
	fmt.Scanln(&config.Paths.SSL)
	fmt.Println("Enter vhost file path: ")
	fmt.Scanln(&config.Paths.Vhost)

	file, _ := json.MarshalIndent(config, "", " ")

	_ = ioutil.WriteFile(configPath+"/config.json", file, 0644)
}
