package main

import (
	"errors"

	"github.com/briansorahan/freesound"
)

// upload uploads a sound
func (f freesnd) upload(args []string) error {
	if err := f.setAccessToken(); err != nil {
		return err
	}
	resp, err := f.c.Upload(args[0], freesound.Upload{})
	if err != nil {
		return err
	}
	if resp.StatusCode < 300 {
		_ = resp.Body.Close() // Best effort.
		return errors.New(resp.Status)
	}
	return resp.Body.Close()
}
