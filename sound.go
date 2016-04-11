package freesound

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// Upload contains optional params for sound upload.
type Upload struct {
	Name        string
	Tags        string
	Description string
	License     string
	Pack        string
	Geotag      string
}

// WriteFields writes upload fields to a multipart writer.
// TODO: support fields other than pack.
func (upload Upload) WriteFields(w *multipart.Writer) error {
	if upload.Pack != "" {
		if err := w.WriteField("pack", upload.Pack); err != nil {
			return err
		}
	}
	return nil
}

// Upload uploads the sound located at pathAudio.
// This requires an oauth access token.
func (c *Client) Upload(pathAudio string, upload Upload) (*http.Response, error) {
	if c.accessToken == "" {
		return nil, ErrMissingToken
	}
	file, err := os.Open(pathAudio)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }() // Best effort.

	var (
		u      = secURL + "/sounds/upload/"
		body   = &bytes.Buffer{}
		writer = multipart.NewWriter(body)
	)
	part, err := writer.CreateFormFile("audiofile", filepath.Base(pathAudio))
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// PendingUploads returns your pending uploads.
func (c *Client) PendingUploads() (*http.Response, error) {
	if c.accessToken == "" {
		return nil, ErrMissingToken
	}
	u := secURL + "/sounds/pending_uploads/"
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, errors.New(resp.Status)
	}
	return resp, nil
}
