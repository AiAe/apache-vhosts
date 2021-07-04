package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
)

func IsFramework(project string) bool {
	_, err := os.Stat(Cfg.Paths.Projects + "\\" + project + "\\public")
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TruncateFile() {
	err := os.Truncate(Cfg.Paths.Vhost, 0)
	if err != nil {
		fmt.Println("Vhost file could not be found, please check your config!")
		panic(err)
	}
}

func SaveToFile(template string) {
	f, err := os.OpenFile(Cfg.Paths.Vhost, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(template)

	if err2 != nil {
		panic(err2)
	}

	defer f.Close()
}

func RunCommand(name string, args []string) {
	command := exec.Command(name, args...)
	_, err := command.Output()
	if err != nil {
		panic(err)
	}
}

func FetchProjects() []string {
	var files []string

	contents, err := ioutil.ReadDir(Cfg.Paths.Projects)

	if err != nil {
		fmt.Println("Projects folder could not be found, please check your config!")
		panic(err)
	}

	for _, f := range contents {
		if f.IsDir() && !Contains(Cfg.Dirs.Ignore, f.Name()) {
			files = append(files, f.Name())
		}
	}

	return files
}

func CreateVhost(project string) {
	path := Cfg.Paths.Projects + "/" + project

	if IsFramework(project) {
		path = path + "/public"
	}

	noSSL := fmt.Sprintf(Cfg.Template.NoSSL, path, project)
	SaveToFile(noSSL)

	if Plt.Darwin {
		CreateSSL(project+".test", project+"-key.pem", project+".pem")
		SSL := fmt.Sprintf(Cfg.Template.SSL, path, project, Cfg.Paths.SSL)
		SaveToFile(SSL)
	}
}

func CreateSSL(host string, key string, cert string) {
	keyFile := Cfg.Paths.SSL + "/" + key
	certFile := Cfg.Paths.SSL + "/" + cert

	RunCommand("mkcert", []string{"-key-file", keyFile, "-cert-file", certFile, host})
}

func userPath() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}
