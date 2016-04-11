package main

import (
	"os"

	"github.com/briansorahan/freesound"
)

func (f freesnd) authorize(c *freesound.Client, args []string) {
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}
	authCode := args[0]
}
