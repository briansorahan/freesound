package main

import (
	"bytes"
	"os"
)

// setAccessToken sets the freesound client access token
// by trying to read it from a file.
func (f freesnd) setAccessToken() error {
	accessTokenFile, err := os.Open(pathAccess)
	if err != nil {
		return err
	}
	accessToken := make([]byte, 128)
	if _, err := accessTokenFile.Read(accessToken); err != nil {
		return err
	}
	accessToken = bytes.TrimSpace(bytes.Trim(accessToken, "\x00"))
	f.client.SetAccessToken(string(accessToken))
	return nil
}
