package main

import (
	"log"
	"os"

	"github.com/briansorahan/freesound"
)

func main() {
	key, secret := os.Getenv("FREESOUND_API_KEY"), os.Getenv("FREESOUND_API_SECRET")
	// TODO: use client
	if _, err := freesound.New(key, secret); err != nil {
		log.Fatal(err)
	}
}
