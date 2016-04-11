package freesound

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User represents a user of the freesound API.
type User struct {
	URL                string `json:"url,omitempty"`
	Name               string `json:"username,omitempty"`
	About              string `json:"about,omitempty"`
	Ref                string `json:"ref,omitempty"`
	HomePage           string `json:"home_page,omitempty"`
	Avatar             Avatar `json:"avatar,omitempty"`
	Joined             Time   `json:"date_joined,omitempty"`
	Sounds             string `json:"sounds,omitempty"`
	Packs              string `json:"packs,omitempty"`
	NumSounds          int    `json:"num_sounds"`
	NumPacks           int    `json:"num_packs"`
	NumComments        int    `json:"num_comments"`
	BookmarkCategories string `json:"bookmark_categories,omitempty"`
}

// Avatar contains the links for a user's avatar.
type Avatar struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

// GetUser gets a user profile.
func (c *Client) GetUser(name string) (*User, error) {
	req, err := http.NewRequest("GET", BaseURL+"/users/"+name+"/", nil)
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
