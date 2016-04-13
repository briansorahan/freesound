package main

import "os"

func (f freesnd) authorize(args []string) error {
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}
	return nil
}
