package main

import (
	"github.com/PaulTaranu/CarTrack/login-service/config"
	"github.com/PaulTaranu/CarTrack/login-service/router"
)

func main() {
	config.InitDB()

	e := router.InitRoutes()
	e.Logger.Fatal(e.Start(":8080"))
}
