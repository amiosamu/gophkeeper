package main

import (
	"log"

	"github.com/amiosamu/gophkeeper/auth-service/internal/app"
)

func main() {
	if err := app.NewApp().Run(); err != nil {
		log.Fatal(err)
	}
}
