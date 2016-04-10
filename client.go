package freesound

import (
	"errors"
	"net/http"
)

// Client represents a freesound API client.
type Client struct {
	Key    string
	Secret string

	httpClient *http.Client
}

func (c *Client) SoundSearch(query SoundSearchQuery) (*SoundSearchResult, error) {
	return nil, errors.New("Not Implemented")
}

// NewClient creates a new freesound API client.
func NewClient(key, secret string) (*Client, error) {
	return &Client{
		Key:        key,
		Secret:     secret,
		httpClient: &http.Client{},
	}, nil
}
