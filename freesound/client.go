package freesound

const (
	V1 = 100
	V2 = 200
)

type Client interface {
	SoundSearch(query string)
}

func NewClient(apiKey string, version int) *Client {
	if (version == V1) {
		return NewClientV1(apiKey)
	} else if (version == V2) {
		return NewClientV2(apiKey)
	} else {
	}
}
