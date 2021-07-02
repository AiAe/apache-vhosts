package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

var cfg Config
var platform Platform

func fetchProjects() []string {
	var files []string

	contents, err := ioutil.ReadDir(cfg.Paths.Projects)

	if err != nil {
		panic(err)
	}

	for _, f := range contents {
		if f.IsDir() && !contains(cfg.Dirs.Ignore, f.Name()) {
			files = append(files, f.Name())
		}
	}

	return files
}

func createVhost(project string) {
	path := cfg.Paths.Projects + "/" + project

	if isFramework(project) {
		path = path + "/public"
	}

	noSSL := fmt.Sprintf(cfg.Template.NoSSL, path, project)
	saveToFile(noSSL)

	if platform.Darwin {
		createSSL(project)
		SSL := fmt.Sprintf(cfg.Template.SSL, path, project, cfg.Paths.SSL)
		saveToFile(SSL)
	}
}

func createSSL(host string) {
	keyFile := cfg.Paths.SSL + "/" + host + "-key.pem"
	certFile := cfg.Paths.SSL + "/" + host + ".pem"
	command := exec.Command("mkcert", "-key-file", keyFile, "-cert-file", certFile, host+".test")
	_, err := command.Output()
	if err != nil {
		panic(err)
	}
}

func restartHttpd() {
	command := exec.Command("brew", "services", "restart", "httpd")
	_, err := command.Output()
	if err != nil {
		panic(err)
	}
}

func createLocalhost() {
	keyFile := cfg.Paths.SSL + "/key.pem"
	certFile := cfg.Paths.SSL + "/cert.pem"
	command := exec.Command("mkcert", "-key-file", keyFile, "-cert-file", certFile, "localhost")
	_, err := command.Output()
	if err != nil {
		panic(err)
	}
}

func main() {
	checkPlatform(&platform)
	readFile(&cfg)
	truncateFile()

	if !platform.Darwin {
		fmt.Println("This platform is not supported, creating ssl and restarting server wont work!")
	} else {
		createLocalhost()
	}

	files := fetchProjects()
	for _, project := range files {
		fmt.Printf("Creating %v.test\n", project)
		createVhost(project)
	}

	if platform.Darwin {
		restartHttpd()
	}
}
