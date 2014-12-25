package freesound

type ClientV1 struct {
	apiKey string
}

func NewClientV1(apiKey string) *Client {
	return &ClientV1{apiKey}
}
