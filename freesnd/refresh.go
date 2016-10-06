package main

import (
	"bytes"
	"os"
)

// refreshToken refreshes an oauth access token.
func (f freesnd) refreshToken(args []string) error {
	// Read the refresh token from a file.
	refreshFile, err := os.Open(pathRefresh)
	if err != nil {
		return err
	}
	refreshToken := make([]byte, 128)
	if _, err := refreshFile.Read(refreshToken); err != nil {
		return err
	}
	refreshToken = bytes.TrimSpace(bytes.Trim(refreshToken, "\x00"))

	// Refresh the access token and write the new one to a file.
	tokenResponse, err := f.client.RefreshAccessToken(string(refreshToken))
	if err != nil {
		return err
	}
	return f.writeTokenResponse(tokenResponse)
}
