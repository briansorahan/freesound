package main

import (
	"fmt"

	"github.com/briansorahan/freesound"
)

// getCode prints the URL that users must visit to generate an auth code.
func getCode(c *freesound.Client, args []string) {
	fmt.Println(c.CodeURL())
}
