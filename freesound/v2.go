package freesound

type ClientV2 struct {
	apiKey string
}

func NewClientV2(apiKey string) *Client {
	return &ClientV2{apiKey}
}
