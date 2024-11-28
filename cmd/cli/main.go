package main

import (
	"asciiartweb/nyeltay/algaliyev/internal"
	"log"
)

func main() {
	if err := internal.Run(); err != nil {
		log.Print(err)
	}
}
