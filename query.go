package freesound

import (
	"errors"
	"net/url"
	"strconv"
)

// addValue adds a key/value pair to a url.Values if value is not
// the empty string
func addValue(values url.Values, key, value string) {
	if value != "" {
		values.Add(key, value)
	}
}

// addIntValue adds a key to a url.Values if value is not 0
// returns an error if the value int could not be converted
// to a string
func addIntValue(values url.Values, key string, value int) {
	if value != 0 {
		val := strconv.Itoa(value)
		values.Add(key, val)
	}
}

func addBoolValue(values url.Values, key string, value bool) {
	if value {
		values.Add(key, "1")
	} else {
		values.Add(key, "0")
	}
}

func getValues(query interface{}) (url.Values, error) {
	values := url.Values{}
	if v, ok := query.(SoundSearchQuery); ok {
		addValue(values, "q", v.Query)
		addIntValue(values, "p", v.Page)
		addValue(values, "f", v.Filter.String())
		addValue(values, "s", v.Sort)
		addValue(values, "fields", v.Fields)
		addIntValue(values, "sounds_per_page", v.SoundsPerPage)
		addBoolValue(values, "g", v.GroupInPacks)
	} else {
		return nil, errors.New("unrecognized query type")
	}
	return values, nil
}
