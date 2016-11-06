package main

import (
	"os"
	"time"

	"github.com/briansorahan/freesound"
)

// authorize stores an auth code, access token, and refresh token.
func (f freesnd) authorize(args []string) error {
	if len(args) < 1 {
		_ = usage(args) // Never returns an error.
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
	now := time.Now()
	tokenResponse, err := f.client.GetAccessToken(args[0])
	if err != nil {
		return err
	}
	return f.writeTokenResponse(tokenResponse, now)
}

// writeTokenResponse writes an access token and refresh token to disk.
// It uses now to write an expires_at timestamp to a file to know when
// the token needs to be refreshed.
func (f freesnd) writeTokenResponse(resp freesound.AccessTokenResponse, now time.Time) error {
	// Write the access token.
	accessTokenFile, err := os.Create(pathAccess)
	if err != nil {
		return err
	}
	if _, err := accessTokenFile.Write([]byte(resp.AccessToken)); err != nil {
		return err
	}

	// Write the refresh token.
	refreshTokenFile, err := os.Create(pathRefresh)
	if err != nil {
		return err
	}
	if _, err = refreshTokenFile.Write([]byte(resp.RefreshToken)); err != nil {
		return err
	}

	// Write the expires_at timestamp.
	expirationFile, err := os.Create(pathExpiration)
	if err != nil {
		return err
	}
	expiration := now.Add(time.Duration(resp.ExpiresIn) * time.Second).UTC()
	_, err = expirationFile.Write([]byte(expiration.Format(time.RFC3339)))
	return err
}
