package main

import (
	"fmt"
	"github.com/AiAe/apache-vhosts/src/utils"
)

func main() {
	utils.ReadFile(&utils.Cfg)
	utils.CheckPlatform(&utils.Plt)
	utils.TruncateFile()

	if !utils.Plt.Darwin {
		fmt.Println("This platform is not supported, creating ssl and restarting server are disabled!")
	}

	for _, project := range utils.FetchProjects() {
		fmt.Printf("Creating %v.test\n", project)
		utils.CreateVhost(project)
	}

	if utils.Plt.Darwin {
		utils.CreateSSL("localhost", "key.pem", "cert.pem")
		utils.RunCommand("brew", []string{"services", "restart", "httpd"})
	}
}
