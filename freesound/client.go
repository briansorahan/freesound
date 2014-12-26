package freesound

import (
	"fmt"
)

const (
	V1 = 100
	V2 = 200
)

func NewClient(apiKey string, version int) (Client, error) {
	if version == V1 {
		return NewClientV1(apiKey)
	} else if version == V2 {
		return NewClientV2(apiKey)
	} else {
		return nil, fmt.Errorf("unrecognized version %d", version)
	}
}
