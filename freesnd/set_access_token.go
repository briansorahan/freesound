package main

import "os"

// setAccessToken sets the freesound client access token
// by trying to read it from a file
func (f freesnd) setAccessToken() error {
	cf, err := os.Open(pathCode)
	if err != nil {
		return err
	}
	code := make([]byte, 128)
	if _, err := cf.Read(code); err != nil {
		return err
	}
	f.c.SetAccessToken(string(code))
	return nil
}
