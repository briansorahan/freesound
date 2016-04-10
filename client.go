package freesound

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "http://www.freesound.org/apiv2"
)

// Client represents a freesound API client.
type Client struct {
	Token string

	httpClient *http.Client
}

// Do performs an http request.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Token "+c.Token)
	return c.httpClient.Do(req)
}

// GetUser gets a user profile.
func (c *Client) GetUser(name string) (*User, error) {
	req, err := http.NewRequest("GET", baseURL+"/users/"+name+"/", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s", resp.Status)
	}
	u := &User{}
	if err := json.NewDecoder(resp.Body).Decode(u); err != nil {
		_ = resp.Body.Close() // Best effort.
		return nil, err
	}
	_ = resp.Body.Close() // Best effort.
	return u, nil
}

// New creates a new freesound API client.
func New(token string) (*Client, error) {
	return &Client{
		Token:      token,
		httpClient: &http.Client{},
	}, nil
}
