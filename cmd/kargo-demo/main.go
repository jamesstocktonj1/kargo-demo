package main

import (
	"log"

	"github.com/jamesstocktonj1/kargo-demo/app"
)

func main() {
	server, err := app.NewRouter()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
