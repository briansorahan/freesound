package main

import "os"

// authorize stores an auth code
func (f freesnd) authorize(args []string) error {
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}
	cf, err := os.Create(pathCode)
	if err != nil {
		return err
	}
	if _, err := cf.Write([]byte(args[0])); err != nil {
		return err
	}
	return nil
}
