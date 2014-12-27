package freesound

import (
	"errors"
	"testing"
)

func TestNewClient(t *testing.T) {
	c1, err := NewClient("API_KEY", V1)
	if err != nil {
		t.Fatal(err)
	}
	if c1.Version() != V1 {
		t.Fatal(errors.New("wrong version"))
	}
	c2, err := NewClient("API_KEY", V2)
	if err != nil {
		t.Fatal(err)
	}
	if c2.Version() != V2 {
		t.Fatal(errors.New("wrong version"))
	}
	_, err = NewClient("API_KEY", 24)
	if err == nil {
		t.Fail()
	}
}
