package freesound

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

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

	req, err := uploadRequest(filepath.Base(pathAudio), file, upload)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

// uploadRequest returns a new *http.Request for a file upload.
func uploadRequest(basepath string, r io.Reader, upload Upload) (*http.Request, error) {
	var (
		u      = secURL + "/sounds/upload/"
		body   = &bytes.Buffer{}
		writer = multipart.NewWriter(body)
	)
	part, err := writer.CreateFormFile("audiofile", basepath)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, r); err != nil {
		return nil, err
	}
	if err := upload.WriteFields(writer); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", u, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

// PendingUploadsResponse represents a response to /sounds/pending_uploads/
type PendingUploadsResponse struct {
	PendingDescription []string            `json:"pending_description"`
	PendingModeration  []PendingModeration `json:"pending_moderation"`
}

// PendingModeration represents an uploaded sound that is pending moderation.
type PendingModeration struct {
	Description string   `json:"description"`
	License     string   `json:"license"`
	Tags        []string `json:"tags"`
	Created     Time     `json:"created"`
	Images      Images   `json:"images"`
	ID          int      `json:"id"`
	Name        string   `json:"name"`
}

// Images contains links to sound images.
type Images struct {
	WaveformL string `json:"waveform_l"`
	WaveformM string `json:"waveform_m"`
	SpectralL string `json:"spectral_l"`
	SpectralM string `json:"spectral_m"`
}

// PendingUploads returns your pending uploads.
func (c *Client) PendingUploads() (PendingUploadsResponse, error) {
	pendingUploadsResponse := PendingUploadsResponse{}

	if c.accessToken == "" {
		return pendingUploadsResponse, ErrMissingToken
	}
	u := secURL + "/sounds/pending_uploads/"
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return pendingUploadsResponse, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return pendingUploadsResponse, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode >= 300 {
		return pendingUploadsResponse, errors.New(resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&pendingUploadsResponse); err != nil {
		return pendingUploadsResponse, err
	}
	return pendingUploadsResponse, nil
}

// Upload contains optional params for sound upload.
type Upload struct {
	Name        string
	Tags        string
	Description string
	License     string
	Pack        string
	Geotag      string
}

// Validate validates the upload data. This just checks that
// if any one of tags, description, or license is provided
// then they all have been provided.
func (upload Upload) Validate() error {
	if upload.Tags == "" && upload.Description == "" && upload.License == "" {
		return nil
	}
	if upload.Tags != "" && upload.Description != "" && upload.License != "" {
		return nil
	}
	return errors.New("tags, description, or license must all be either empty or not empty")
}

// WriteFields writes upload fields to a multipart writer.
func (upload Upload) WriteFields(w *multipart.Writer) error {
	if err := upload.writeField("name", upload.Name, w); err != nil {
		return err
	}
	if err := upload.writeField("tags", upload.Tags, w); err != nil {
		return err
	}
	if err := upload.writeField("description", upload.Description, w); err != nil {
		return err
	}
	if err := upload.writeField("license", upload.License, w); err != nil {
		return err
	}
	if err := upload.writeField("pack", upload.Pack, w); err != nil {
		return err
	}
	return upload.writeField("geotag", upload.Geotag, w)
}

func (upload Upload) writeField(key, value string, w *multipart.Writer) error {
	if value != "" {
		return w.WriteField(key, value)
	}
	return nil
}
