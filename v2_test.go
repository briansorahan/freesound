package freesound

import (
	"testing"
)

func TestNewClientV2(t *testing.T) {
	_, err := NewClientV2("API_KEY")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSoundSearchV2(t *testing.T) {
	c, err := NewClientV2("API_KEY")
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.SoundSearch(SoundSearchQuery{Query:"dog bark"})
	if err == nil {
		t.Fail()
	}
}
