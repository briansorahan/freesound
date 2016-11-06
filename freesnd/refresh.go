package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"time"
)

// refreshToken refreshes an oauth access token.
func (f freesnd) refreshToken(args []string) error {
	refreshToken, err := readRefreshToken()
	if err != nil {
		return err
	}
	// Refresh the access token and write the new one to a file.
	now := time.Now()
	tokenResponse, err := f.client.RefreshAccessToken(string(refreshToken))
	if err != nil {
		return err
	}
	return f.writeTokenResponse(tokenResponse, now)
}

// readRefreshToken reads the refresh token.
func readRefreshToken() (string, error) {
	// Read the refresh token from a file.
	refreshFile, err := os.Open(pathRefresh)
	if err != nil {
		return "", err
	}
	refreshToken := make([]byte, 128)
	if _, err := refreshFile.Read(refreshToken); err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(bytes.Trim(refreshToken, "\x00"))), nil
}

// freshToken ensures the access token is not expired.
func (f freesnd) freshToken(args []string) error {
	// Read the expires_at timestamp.
	expiresAt, err := readExpiration()
	if err != nil {
		return err
	}

	// If the access token has not expired just run cmdfunc,
	// otherwise get a new access token and then run cmdfunc.
	if time.Now().UTC().Before(expiresAt) {
		return nil
	}
	return f.refreshToken(args)
}

// readExpiration reads the expires_at file.
func readExpiration() (time.Time, error) {
	expirationFile, err := os.Open(pathExpiration)
	if err != nil {
		return time.Time{}, err
	}
	buf, err := ioutil.ReadAll(expirationFile)
	if err != nil {
		return time.Time{}, err
	}
	return time.Parse(string(buf), time.RFC3339)
}
