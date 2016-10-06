package main

import (
	"os"

	"github.com/briansorahan/freesound"
)

// authorize stores an auth code, access token, and refresh token.
func (f freesnd) authorize(args []string) error {
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}

	// Write the auth code to a file.
	codeFile, err := os.Create(pathCode)
	if err != nil {
		return err
	}
	if _, err := codeFile.Write([]byte(args[0])); err != nil {
		return err
	}

	// Get an access token using the auth code and write it and
	// the refresh token to a file.
	tokenResponse, err := f.client.GetAccessToken(args[0])
	if err != nil {
		return err
	}
	return f.writeTokenResponse(tokenResponse)
}

// writeTokenResponse writes an access token and refresh token to disk.
func (f freesnd) writeTokenResponse(resp freesound.AccessTokenResponse) error {
	accessTokenFile, err := os.Create(pathAccess)
	if err != nil {
		return err
	}
	if _, err := accessTokenFile.Write([]byte(resp.AccessToken)); err != nil {
		return err
	}
	refreshTokenFile, err := os.Create(pathRefresh)
	if err != nil {
		return err
	}
	_, err = refreshTokenFile.Write([]byte(resp.RefreshToken))
	return err
}
