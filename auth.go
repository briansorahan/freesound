package freesound

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var secURL = strings.Replace(BaseURL, "http", "https", 1)

// CodeURL
func (c *Client) CodeURL() string {
	values := url.Values{}
	values.Set("client_id", c.ID)

	return fmt.Sprintf("%s/oauth2/authorize?client_id=%s&response_type=code", secURL, c.ID)
}

// GetAccessToken gets an oauth access token with the provided auth code.
func (c *Client) GetAccessToken(authCode string) (*http.Response, error) {
	// Set up the query params.
	values := url.Values{}
	values.Set("client_id", c.ID)
	values.Set("client_secret", c.Secret)
	values.Set("grant_type", "authorization_code")
	values.Set("code", authCode)

	// Make the request.
	u, err := url.Parse(secURL + "/oauth2/access_token")
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.PostForm(u.String(), values)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 {
		return nil, errors.New(resp.Status)
	}
	return resp, nil
}

// SetAccessToken sets the access token to be used by the client.
func (c *Client) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
}
