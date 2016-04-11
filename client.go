package freesound

import (
	"errors"
	"net/http"
)

const (
	// BaseURL is the base URL for the freesound API v2.
	BaseURL = "http://www.freesound.org/apiv2"
)

var (
	// ErrMissingToken is returned whenever you call a method
	// that requires an access token before you call
	// SetAccessToken with a valid access token.
	ErrMissingToken = errors.New("missing access token")
)

// Client represents a freesound API client.
type Client struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`

	accessToken string
	httpClient  *http.Client
}

// Do performs an http request.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, errors.New(resp.Status)
	}
	return resp, nil
}

// New creates a new freesound API client.
func New(id, secret string) (*Client, error) {
	return &Client{
		ID:     id,
		Secret: secret,

		httpClient: &http.Client{},
	}, nil
}
