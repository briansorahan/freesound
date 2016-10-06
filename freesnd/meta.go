package main

import (
	"fmt"
	"os"

	"github.com/briansorahan/sndtag"
)

// meta prints metadata about all the files listed in args.
func (f freesnd) meta(args []string) error {
	for _, filename := range args {
		fd, err := os.Open(filename)
		if err != nil {
			return nil
		}
		defer func() { _ = fd.Close() }() // Best effort.

		tags, err := sndtag.New(fd)
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", filename)
		for k, v := range tags {
			fmt.Printf("\t%-20s: %s\n", k, v)
		}
	}
	return nil
}
