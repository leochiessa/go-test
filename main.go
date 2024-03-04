package main

import (
	"go-test/pkg/containers"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	var app containers.AppContainer = containers.NewAppContainer()
	app.ServerContainer.Server.Run()
}
