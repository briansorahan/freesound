package freesound

import (
	"net/url"
	"os"
	"testing"
)

const (
	apiKeyVar = "FREESOUND_API_KEY"
	baseUrl   = "http://www.freesound.org/api"
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
	_, err := NewClientV1("MY_API_KEY")
	if err != nil {
		t.Fail()
	}
}

func TestClientUrl(t *testing.T) {
	c, err := NewClientV1("MY_API_KEY")
	if err != nil {
		t.Fail()
	}
	values := url.Values{}
	values.Add("foo", "bar baz")
	client := c.(*ClientV1)
	u := client.Url("sounds/search", values)
	if u != baseUrl+"/sounds/search?api_key=MY_API_KEY&foo=bar+baz" {
		t.Fail()
	}
}

func TestSoundSearchV1(t *testing.T) {
	c := getClient(t)
	_, err := c.SoundSearch("cat meow")
	if err != nil {
		t.Fatal(err)
	}
	bc, err := NewClientV1("FOOBAR")
	if err != nil {
		t.Fatal(err)
	}
	_, err = bc.SoundSearch("hinge creak")
	if err == nil {
		t.Fail()
	}
}
