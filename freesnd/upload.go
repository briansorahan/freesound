package main

import (
	"encoding/json"
	"errors"
	"flag"
	"os"

	"github.com/briansorahan/freesound"
)

// upload uploads a sound.
func (f freesnd) upload(args []string) error {
	if err := f.setAccessToken(); err != nil {
		return err
	}

	// Parse the upload options.
	upload, remainingArgs, err := getUpload(args)
	if err != nil {
		return err
	}

	// TODO: handle multiple uploads
	resp, err := f.client.Upload(remainingArgs[0], upload)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		_ = resp.Body.Close() // Best effort.
		return errors.New(resp.Status)
	}
	return resp.Body.Close()
}

// getUpload returns an upload struct with fields set from command-line options.
// It also returns the non-flag command-line arguments and an error.
func getUpload(args []string) (freesound.Upload, []string, error) {
	var (
		upload = freesound.Upload{}
		fs     = flag.NewFlagSet("upload", cliErrorHandling)
	)
	fs.StringVar(&upload.Tags, "tags", "", "Space-separated tags.")
	fs.StringVar(&upload.Description, "description", "", "Description of the sound.")
	fs.StringVar(&upload.License, "license", "", "License for the sound. Must be one of 'Attribution', 'Attribution Noncommercial', or 'Creative Commons C0'.")
	fs.StringVar(&upload.Pack, "pack", "", "Pack for the new sound. If one doesn't exist with the provided name, a new one will be created.")
	fs.StringVar(&upload.Geotag, "geotag", "", "Geotag for the new sound. Should be of the form 'lat,long,zoom' (e.g. '2.145677,3.22345,14').")
	if err := fs.Parse(args); err != nil {
		return upload, nil, err
	}
	return upload, fs.Args(), upload.Validate()
}

// pendingUploads print JSON-encoded info about pending uploads.
func (f freesnd) pendingUploads(args []string) error {
	if err := f.setAccessToken(); err != nil {
		return err
	}

	pendingUploads, err := f.client.PendingUploads()
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(pendingUploads)
}
