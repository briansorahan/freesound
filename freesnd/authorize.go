package main

import (
	"os"
)

func authorize() int {
	if os.IsExist(tokenFile) {
		f, err := os.Open(tokenFile)
		if err != nil {
		}
	}
	return 0
}
