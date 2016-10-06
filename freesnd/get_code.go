package main

import "fmt"

// getCode prints the URL that users must visit to generate an auth code.
func (f freesnd) getCode(args []string) error {
	_, err := fmt.Println(f.client.CodeURL())
	return err
}
