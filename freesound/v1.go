package freesound

type ClientV1 struct {
	apiKey string
}

func (c *ClientV1) SoundSearch(query string) (SoundSearchResult, error) {
}

func NewClientV1(apiKey string) (Client, error) {
	c := ClientV1{apiKey}
	return &c, nil
}
