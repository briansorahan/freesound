package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

var (
	pathHome   = path.Join(os.Getenv("HOME"), ".freesnd")
	pathKey    = path.Join(pathHome, "key")
	pathSecret = path.Join(pathHome, "secret")
	pathAccess = path.Join(pathHome, "access")
)

func usage() {
	fmt.Fprintf(os.Stderr, "%s COMMAND [OPTIONS]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "COMMANDS\n")
	fmt.Fprintf(os.Stderr, "authorize CODE           Fetch an access token and store locally.\n")
	fmt.Fprintf(os.Stderr, "get-code                 Show a URL that allows you to generate an authorization code.\n")
}

func main() {
	// Check arguments.
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	// Ensure the .freesnd directory exists.
	if err := makeHome(); err != nil {
		log.Fatal(err)
	}
	// Commands that can run without a key and secret.
	if os.Args[1] == "configure" {
		if err := configure(os.Args[2:]); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
	// Get the key and secret.
	key, secret, err := getKeySecret()
	if err != nil {
		log.Fatal(err)
	}
	// Initialize client.
	f, err := newFreesnd(key, secret)
	if err != nil {
		log.Fatal(err)
	}
	switch os.Args[1] {
	case "authorize":
		f.authorize(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Unrecognized command: %s\n", os.Args[1])
		usage()
	}
}
