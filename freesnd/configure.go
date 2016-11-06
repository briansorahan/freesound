package main

import (
	"flag"
	"os"
)

// configure saves the key and secret.
func configure(args []string) error {
	// Get key and secret from CLI.
	var (
		fs     = flag.NewFlagSet("configure", flag.ExitOnError)
		key    = fs.String("key", "", "application key (ID)")
		secret = fs.String("secret", "", "application secret")
	)
	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			_ = usage(args) // Never returns an error.
			os.Exit(1)
		}
		return err
	}
	// Write key.
	keyFile, err := os.Create(pathKey)
	if err != nil {
		return err
	}
	if _, err := keyFile.Write([]byte(*key)); err != nil {
		return err
	}
	// Write secret.
	secretFile, err := os.Create(pathSecret)
	if err != nil {
		return err
	}
	if _, err := secretFile.Write([]byte(*secret)); err != nil {
		return err
	}
	return nil
}
