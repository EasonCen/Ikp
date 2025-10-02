package main

import (
	"Ikp_gin/config"
	"Ikp_gin/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.AppConfig.App.Port
	r.Run(port)
}
