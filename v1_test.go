package freesound

import (
	"net/url"
	"os"
	"testing"
)

const (
	apiKeyVar = "FREESOUND_API_KEY"
	baseURL   = "http://www.freesound.org/api"
)

func getClient(t *testing.T) Client {
	k := os.Getenv(apiKeyVar)
	c, err := NewClientV1(k)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestNewClientV1(t *testing.T) {
	if _, err := NewClientV1("MY_API_KEY"); err != nil {
		t.Fail()
	}
}

func TestClientURL(t *testing.T) {
	c, err := NewClientV1("MY_API_KEY")
	if err != nil {
		t.Fail()
	}
	values := url.Values{}
	values.Add("foo", "bar baz")
	client := c.(*ClientV1)
	u := client.URL("sounds/search", values)
	if u != baseURL+"/sounds/search?api_key=MY_API_KEY&foo=bar+baz" {
		t.Fail()
	}
}

func TestSoundSearchV1(t *testing.T) {
	c := getClient(t)
	if _, err := c.SoundSearch(SoundSearchQuery{Query: "cat meow"}); err != nil {
		t.Fatal(err)
	}
	bc, err := NewClientV1("FOOBAR")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = bc.SoundSearch(SoundSearchQuery{Query: "hinge creak"}); err == nil {
		t.Fail()
	}
}
