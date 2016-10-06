package freesound

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
)

const (
	// BaseURL is the base URL for the freesound API v2.
	BaseURL = "http://www.freesound.org/apiv2"
)

// Verbosity levels.
const (
	VerbosityOff = iota
	VerbosityOn
)

var (
	// ErrMissingToken is returned whenever you call a method
	// that requires an access token before you call
	// SetAccessToken with a valid access token.
	ErrMissingToken = errors.New("missing access token")
)

// Client represents a freesound API client.
type Client struct {
	ID        string `json:"id"`
	Secret    string `json:"secret"`
	Verbosity int    `json:"verbosity"`

	accessToken string
	httpClient  *http.Client
}

// Do performs an http request.
// Callers are expected to close the response body when the returned error is nil.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+c.accessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		errmsg, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			_ = resp.Body.Close() // Best effort.
			return nil, err
		}
		_ = resp.Body.Close() // Best effort.
		return nil, errors.New(string(errmsg))
	}
	return resp, nil
}

// logRequest logs an http request.
func (c *Client) logRequest(req *http.Request) error {
	if c.Verbosity == VerbosityOff {
		return nil
	}
	reqDump, err := httputil.DumpRequestOut(req, false)
	if err != nil {
		return err
	}
	_, err = os.Stderr.Write(reqDump)
	return err
}

// New creates a new freesound API client.
func New(id, secret string) (*Client, error) {
	return &Client{
		ID:     id,
		Secret: secret,

		httpClient: &http.Client{},
	}, nil
}
