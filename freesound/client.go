package freesound

import (
	"fmt"
)

const (
	V1 = 100
	V2 = 200
)

type User struct {
	Username string `json:"username,omitempty"`
	URL string `json:"url,omitempty"`
	Ref string `json:"ref,omitempty"`
}

type SoundSearchResult struct {
	URL string `json:"url,omitempty"`
}

type Client interface {
	SoundSearch(query string) (SoundSearchResult, error)
}

func NewClient(apiKey string, version int) (Client, error) {
	if (version == V1) {
		return NewClientV1(apiKey)
	} else if (version == V2) {
		return NewClientV2(apiKey)
	} else {
		return nil, fmt.Errorf("unrecognized version %d", version)
	}
}
