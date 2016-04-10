package freesound

import (
	"fmt"
)

// Version is an API version.
type Version int

const (
	// V1 is API version 1.
	V1 = Version(1)
	// V2 is API version 2.
	V2 = Version(2)
)

// NewClient creates a new freesound API client.
func NewClient(apiKey string, version Version) (Client, error) {
	if version == V1 {
		return newClientV1(apiKey)
	} else if version == V2 {
		return newClientV2(apiKey)
	} else {
		return nil, fmt.Errorf("unrecognized version %d", version)
	}
}
