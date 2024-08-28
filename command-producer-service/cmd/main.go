package main

import (
	"log"

	"github.com/amiosamu/gophkeeper/command-producer-service/internal/app"
)

func main() {
	if err := app.NewApp().Run(); err != nil {
		log.Fatal(err)
	}
}
