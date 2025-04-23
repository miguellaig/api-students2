package main

import (
	"log"

	"github.com/miguellaig/api-students2/api"
)

func main() {

	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
