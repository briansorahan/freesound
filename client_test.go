package freesound

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient(os.Getenv("FREESOUND_API_TOKEN"))
	if err != nil {
		t.Fatal(err)
	}
	u, err := c.GetUser("wjoojoo")
	if err != nil {
		t.Fatal(err)
	}
	if expected, got := "wjoojoo", u.Name; expected != got {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}
