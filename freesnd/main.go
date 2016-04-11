package main

import (
	"fmt"
	"log"
	"os"
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

	// Initialize client.
	f, err := newFreesnd()
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "authorize":
		f.authorize(os.Args[2:])
	case "configure":
		f.configure(os.Args[2:])
	case "get-code":
		f.getCode(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Unrecognized command: %s\n", os.Args[1])
		usage()
	}
}
