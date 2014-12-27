package freesound

import (
	"encoding/json"
	"errors"
	"net/http"
)

// apiError decodes JSON response bodies and returns a useful error
func apiError(method, url string, response *http.Response) error {
	err := new(ApiError)
	dec := json.NewDecoder(response.Body)
	ed := dec.Decode(err)
	if ed != nil {
		return ed
	}
	return errors.New(err.Explanation)
}
