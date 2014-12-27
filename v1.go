package freesound

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	SchemeV1   = "http"
	HostV1     = "www.freesound.org"
	BasePathV1 = "api"
)

type ClientV1 struct {
	apiKey     string
	httpClient *http.Client
}

func (c *ClientV1) Url(path string, values url.Values) string {
	values.Add("api_key", c.apiKey)
	u := url.URL{
		Scheme:   SchemeV1,
		Host:     HostV1,
		Path:     strings.Join([]string{BasePathV1, path}, "/"),
		RawQuery: values.Encode(),
	}
	return u.String()
}

func (c *ClientV1) Version() int {
	return V1
}

func (c *ClientV1) SoundSearch(query string) (*SoundSearchResult, error) {
	values := url.Values{}
	values.Add("q", query)
	request, err := http.NewRequest("GET", c.Url("sounds/search", values), nil)
	if err != nil {
		return nil, err
	}
	r, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(r.Body)
	response := new(SoundSearchResult)
	err = dec.Decode(response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewClientV1(apiKey string) (Client, error) {
	c := ClientV1{
		apiKey,
		http.DefaultClient,
	}
	return &c, nil
}
