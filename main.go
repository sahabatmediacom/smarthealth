package main

import (
	"pamer-api/config"
	"pamer-api/router"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	r := router.InitializeRouter()

	r.Run()
}
