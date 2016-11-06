package freesound

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var secURL = strings.Replace(BaseURL, "http", "https", 1)

// CodeURL returns a url used to create an authorization code.
func (c *Client) CodeURL() string {
	values := url.Values{}
	values.Set("client_id", c.ID)
	return fmt.Sprintf("%s/oauth2/authorize?client_id=%s&response_type=code", secURL, c.ID)
}

// AccessTokenResponse is the structure of a response from
// the /apiv2/oauth2/access_token/ endpoint.
type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int64  `json:"expires_in"` // seconds
	RefreshToken string `json:"refresh_token"`
}

// GetAccessToken gets an oauth access token with the provided auth code.
// Client code is responsible for closing the response body.
func (c *Client) GetAccessToken(authCode string) (AccessTokenResponse, error) {
	return c.getAccessToken(authCode, "authorization_code")
}

// RefreshAccessToken refreshes an access token that has expired.
// Access tokens only last 24 hours, so this is a common operation.
func (c *Client) RefreshAccessToken(refreshToken string) (AccessTokenResponse, error) {
	return c.getAccessToken(refreshToken, "refresh_token")
}

// Grant types supported by freesound oauth.
const (
	GrantTypeAuthCode     = "authorization_code"
	GrantTypeRefreshToken = "refresh_token"
)

func (c *Client) getAccessToken(code, grantType string) (AccessTokenResponse, error) {
	// Set up the query params.
	var (
		tokenResponse = AccessTokenResponse{}
		values        = url.Values{}
	)

	// Set a query param based on the grant type.
	switch grantType {
	default:
		return tokenResponse, fmt.Errorf("unsupported grant_type: %s", grantType)
	case GrantTypeAuthCode:
		values.Set("code", code)
	case GrantTypeRefreshToken:
		values.Set("refresh_token", code)
	}

	// Set the rest of the query params.
	values.Set("client_id", c.ID)
	values.Set("client_secret", c.Secret)
	values.Set("grant_type", grantType)

	// Make the request.
	resp, err := c.httpClient.PostForm(secURL+"/oauth2/access_token/", values)
	if err != nil {
		return tokenResponse, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode >= 300 {
		return tokenResponse, errors.New(resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return tokenResponse, err
	}
	return tokenResponse, nil
}

// SetAccessToken sets the access token to be used by the client.
func (c *Client) SetAccessToken(accessToken string) {
	c.accessToken = accessToken
}
