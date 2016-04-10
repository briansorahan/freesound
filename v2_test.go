package freesound

import (
	"testing"
)

func TestNewClientV2(t *testing.T) {
	if _, err := NewClient("API_KEY", "API_SECRET"); err != nil {
		t.Fatal(err)
	}
}
