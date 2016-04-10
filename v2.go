package freesound

import (
	"errors"
	"net/http"
)

type clientV2 struct {
	apiKey     string
	httpClient *http.Client
}

func (c *clientV2) SoundSearch(query SoundSearchQuery) (*SoundSearchResult, error) {
	return nil, errors.New("Not Implemented")
}

func (c *clientV2) Version() Version {
	return V2
}

func newClientV2(apiKey string) (Client, error) {
	c := clientV2{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
	return &c, nil
}
