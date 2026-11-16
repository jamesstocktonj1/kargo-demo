package main

import (
	"log"

	"github.com/jamesstocktonj1/kargo-demo/app"
)

func main() {
	if err := app.NewRouter().Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
