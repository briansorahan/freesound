package freesound

import (
	"errors"
)

type ClientV2 struct {
	apiKey string
}

func (c *ClientV2) SoundSearch(query string) (*SoundSearchResult, error) {
	return nil, errors.New("Not Implemented")
}

func NewClientV2(apiKey string) (Client, error) {
	c := ClientV2{apiKey}
	return &c, nil
}
