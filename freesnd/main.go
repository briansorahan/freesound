package main

import (
	"log"
	"os"

	"github.com/briansorahan/freesound"
)

func main() {
	key, secret := os.Getenv("FREESOUND_API_KEY"), os.Getenv("FREESOUND_API_SECRET")
	c, err := freesound.New(key, secret)
	if err != nil {
		log.Fatal(err)
	}
}
