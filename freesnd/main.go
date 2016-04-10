package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/briansorahan/freesound"
)

var tokenFile = path.Join(os.Getenv("HOME"), ".freesnd_token")

func usage() {
	fmt.Fprintf(os.Stderr, "%s COMMAND [OPTIONS]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "COMMANDS\n")
	fmt.Fprintf(os.Stderr, "authorize CODE           Fetch an access token.\n")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	if os.Args[1] == "authorize" {
		os.Exit(authorize())
	}
}
