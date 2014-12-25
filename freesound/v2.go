package freesound

type ClientV2 struct {
	apiKey string
}

func NewClientV2(apiKey string) (Client, error) {
	c := ClientV2{apiKey}
	return &c, nil
}
