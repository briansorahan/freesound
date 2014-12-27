package freesound

import (
	"errors"
)

type ClientV2 struct {
	apiKey string
}

func (c *ClientV2) SoundSearch(query SoundSearchQuery) (*SoundSearchResult, error) {
	return nil, errors.New("Not Implemented")
}

func (c *ClientV2) Version() int {
	return V2
}

func NewClientV2(apiKey string) (Client, error) {
	c := ClientV2{apiKey}
	return &c, nil
}
