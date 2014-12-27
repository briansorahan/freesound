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

func (c *ClientV1) SoundSearch(query SoundSearchQuery) (*SoundSearchResult, error) {
	const method string = "GET"
	const path string = "sounds/search"

	values, err := getValues(query)
	if err != nil {
		return nil, err
	}
	loc := c.Url(path, values)
	request, err := http.NewRequest(method, loc, nil)
	if err != nil {
		return nil, err
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, apiError(method, loc, response)
	}
	dec := json.NewDecoder(response.Body)
	results := new(SoundSearchResult)
	err = dec.Decode(results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func NewClientV1(apiKey string) (Client, error) {
	c := ClientV1{
		apiKey,
		http.DefaultClient,
	}
	return &c, nil
}
