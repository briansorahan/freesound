package main

import (
	"bytes"
	"errors"
	"os"
)

var (
	ErrEmptyKey    = errors.New("key is empty")
	ErrEmptySecret = errors.New("secret is empty")
)

// getKeySecret gets the key and secret.
func getKeySecret() (string, string, error) {
	// read key
	keyFile, err := os.Open(pathKey)
	if err != nil {
		return "", "", err
	}
	key := make([]byte, 128)
	if _, err := keyFile.Read(key); err != nil {
		return "", "", err
	}
	if string(key) == "" {
		return "", "", ErrEmptyKey
	}

	// read secret
	secret := make([]byte, 256)
	secretFile, err := os.Open(pathSecret)
	if err != nil {
		return "", "", err
	}
	if _, err := secretFile.Read(secret); err != nil {
		return "", "", err
	}
	if string(secret) == "" {
		return "", "", ErrEmptySecret
	}

	// strip trailing null bytes and whitespace
	key = bytes.Trim(bytes.TrimSpace(key), "\x00")
	secret = bytes.Trim(bytes.TrimSpace(secret), "\x00")

	return string(key), string(secret), nil
}
