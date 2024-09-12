package main

import (
	"InfSec/internal/app"
	"log"
)

func main() {
	// прикол
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
