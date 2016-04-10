package freesound

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	schemeV1   = "http"
	hostV1     = "www.freesound.org"
	basePathV1 = "api"
)

type clientV1 struct {
	apiKey     string
	httpClient *http.Client
}

// URL return a url used to query the freesound v1 API
func (c *clientV1) URL(path string, values url.Values) string {
	values.Add("api_key", c.apiKey)
	u := url.URL{
		Scheme:   schemeV1,
		Host:     hostV1,
		Path:     strings.Join([]string{basePathV1, path}, "/"),
		RawQuery: values.Encode(),
	}
	return u.String()
}

// Version return the client version
func (c *clientV1) Version() Version {
	return V1
}

// SoundSearch query the freesound v1 API sound search resource
// see http://www.freesound.org/docs/api/resources_apiv1.html#sound-search-resource
func (c *clientV1) SoundSearch(query SoundSearchQuery) (*SoundSearchResult, error) {
	const method string = "GET"
	const path string = "sounds/search"

	values, err := getValues(query)
	if err != nil {
		return nil, err
	}
	loc := c.URL(path, values)
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

// newclientV1 initialize a new freesound v1 API client
func newClientV1(apiKey string) (Client, error) {
	c := clientV1{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
	return &c, nil
}
