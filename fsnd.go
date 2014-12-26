package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	API_KEY_VAR = "FREESOUND_API_KEY"
	BASE_URL = "http://www.freesound.org/api"
)

func main() {
	key := os.Getenv(API_KEY_VAR)
	requrl := BASE_URL + "/sounds/search?q=contact+mic&api_key=" + key
	log.Println("getting " + requrl)
	resp, err := http.Get(requrl)
	if err != nil {
		log.Println("could not GET " + requrl)
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("could not read response body")
		log.Fatal(err)
	}
	log.Println(bytes.NewBuffer(body).String())
}
