package freesound

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCodeURL(t *testing.T) {
	c, err := New("foo", "bar")
	if err != nil {
		t.Fatal(err)
	}
	if expected, got := "https://www.freesound.org/apiv2/oauth2/authorize?client_id=foo&response_type=code", c.CodeURL(); expected != got {
		t.Fatalf("expected %s, got %s", expected, got)
	}
}

func TestGetAccessToken(t *testing.T) {
	t.SkipNow()

	key, secret := os.Getenv("FREESOUND_API_KEY"), os.Getenv("FREESOUND_API_SECRET")
	c, err := New(key, secret)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.GetAccessToken("e81f05a8cc46c8574405d25545422fe3ef6e1ad3"); err != nil {
		t.Fatal(err)
	}
}

func TestPendingUploads(t *testing.T) {
	t.SkipNow()

	key, secret := os.Getenv("FREESOUND_API_KEY"), os.Getenv("FREESOUND_API_SECRET")
	c, err := New(key, secret)
	if err != nil {
		t.Fatal(err)
	}
	c.SetAccessToken("eed46d7786480c76ee600a3ffaad526cbc295915")
	// resp, err := c.PendingUploads()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// bs, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Println(string(bs))
}

func TestUpload(t *testing.T) {
	const audioFile = "/mnt/2TB102211/freesounds/wjoojoo/minibrute/blip01.wav"

	// Skip the test if the audio file does not exist.
	if _, err := os.Open(audioFile); os.IsNotExist(err) {
		t.SkipNow()
	}

	key, secret := os.Getenv("FREESOUND_API_KEY"), os.Getenv("FREESOUND_API_SECRET")
	c, err := New(key, secret)
	if err != nil {
		t.Fatal(err)
	}
	c.SetAccessToken("eed46d7786480c76ee600a3ffaad526cbc295915")

	resp, err := c.Upload(audioFile, Upload{Pack: "minibrute"})
	if err != nil {
		t.Fatal(err)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bs))
}
