package freesound

import (
	"testing"
)

func TestNewClientV2(t *testing.T) {
	if _, err := newClientV2("API_KEY"); err != nil {
		t.Fatal(err)
	}
}

func TestSoundSearchV2(t *testing.T) {
	c, err := newClientV2("API_KEY")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = c.SoundSearch(SoundSearchQuery{Query: "dog bark"}); err == nil {
		t.Fail()
	}
}
