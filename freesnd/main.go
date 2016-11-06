package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

// cliErrorHandling sets the behavior for errors
// encountered when parsing CLI args.
const cliErrorHandling = flag.ExitOnError

var (
	pathHome       = path.Join(os.Getenv("HOME"), ".freesnd")
	pathKey        = path.Join(pathHome, "key")
	pathSecret     = path.Join(pathHome, "secret")
	pathAccess     = path.Join(pathHome, "access_token")
	pathRefresh    = path.Join(pathHome, "refresh_token")
	pathCode       = path.Join(pathHome, "code")
	pathExpiration = path.Join(pathHome, "expires_at")
)

func usage(args []string) error {
	fmt.Fprintf(os.Stderr, "%s COMMAND [OPTIONS]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "COMMANDS\n")
	fmt.Fprintf(os.Stderr, "authorize CODE           Fetch an access token and store locally.\n")
	fmt.Fprintf(os.Stderr, "get-code                 Show a URL that allows you to generate an authorization code.\n")
	fmt.Fprintf(os.Stderr, "meta FILE [FILE ...]     Display metadata for the given files.\n")
	fmt.Fprintf(os.Stderr, "pending-uploads          Show your pending uploads.\n")
	fmt.Fprintf(os.Stderr, "refresh                  Refresh the access token.\n")
	fmt.Fprintf(os.Stderr, "upload FILE [FILE ...]   Upload a file (requires oauth).\n")
	return nil
}

func main() {
	// Check arguments.
	if len(os.Args) < 2 {
		usage(os.Args)
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
	app, err := newFreesnd(key, secret)
	if err != nil {
		log.Fatal(err)
	}

	// Run the command
	cf, supportedCommand := app.commands[os.Args[1]]
	if !supportedCommand {
		fmt.Fprintf(os.Stderr, "Unrecognized command: %s\n", os.Args[1])
		usage(os.Args)
	}
	if err := cf(os.Args[2:]); err != nil {
		log.Fatal(err)
	}
}
