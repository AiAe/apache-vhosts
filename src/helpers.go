package main

import "os"

func isFramework(project string) bool {
	_, err := os.Stat(cfg.Paths.Projects + "\\" + project + "\\public")
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func truncateFile() {
	err := os.Truncate(cfg.Paths.Vhost, 0)
	if err != nil {
		panic(err)
	}
}

func saveToFile(template string) {
	f, err := os.OpenFile(cfg.Paths.Vhost, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
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
