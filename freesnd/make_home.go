package main

import (
	"os"
	"strings"
)

// makeHome makes the .freesnd directory
func makeHome() error {
	// TODO: better permissions
	if err := os.Mkdir(pathHome, os.ModeDir|os.ModePerm); err != nil {
		if err == os.ErrExist || strings.Contains(err.Error(), "file exists") {
			return nil
		}
		return err
	}
	return nil
}
