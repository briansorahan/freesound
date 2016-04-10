package freesound

import (
	"errors"
	"net/http"
)

// Client represents a freesound API client.
type Client struct {
	Token string

	httpClient *http.Client
}

// SoundSearch searches for sounds.
func (c *Client) SoundSearch(query SoundSearchQuery) (*SoundSearchResult, error) {
	return nil, errors.New("Not Implemented")
}

// NewClient creates a new freesound API client.
func NewClient(token string) (*Client, error) {
	return &Client{
		Token:      token,
		httpClient: &http.Client{},
	}, nil
}
