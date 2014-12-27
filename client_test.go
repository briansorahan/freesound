package freesound

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("API_KEY", V1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = NewClient("API_KEY", V2)
	if err != nil {
		t.Fatal(err)
	}
	_, err = NewClient("API_KEY", 24)
	if err == nil {
		t.Fail()
	}
}
