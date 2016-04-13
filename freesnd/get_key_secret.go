package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrEmptyKey    = errors.New("key is empty")
	ErrEmptySecret = errors.New("secret is empty")
)

// getKeySecret gets the key and secret.
func getKeySecret() (string, string, error) {
	// read key
	key := []byte{}
	keyFile, err := os.Open(pathKey)
	if err != nil {
		return "", "", err
	}
	nr, err := keyFile.Read(key)
	if err != nil {
		return "", "", err
	}
	fmt.Printf("read %d bytes\n", nr)
	if string(key) == "" {
		return "", "", ErrEmptyKey
	}
	// read secret
	secret := []byte{}
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
	return string(key), string(secret), nil
}
