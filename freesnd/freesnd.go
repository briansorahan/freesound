package main

import (
	"os"

	"github.com/briansorahan/freesound"
)

// commandFunc is a func that can run a command.
type commandFunc func([]string) error

// freesnd contains all the state of the program.
type freesnd struct {
	client   *freesound.Client
	home     *os.File
	commands map[string]commandFunc
}

// newFreesnd creates a new freesnd instance.
func newFreesnd(key, secret string) (*freesnd, error) {
	// Initialize the freesound client.
	client, err := freesound.New(key, secret)
	if err != nil {
		return nil, err
	}
	// Check if home is a directory, if it isn't then remove it
	// and recreate it as a directory.
	home, err := os.Open(pathHome)
	if err != nil {
		return nil, err
	}
	homeStat, err := home.Stat()
	if err != nil {
		return nil, err
	}
	if !homeStat.IsDir() {
		if err := os.Remove(pathHome); err != nil {
			return nil, err
		}
		if err := os.Mkdir(pathHome, os.ModeDir); err != nil {
			return nil, err
		}
	}
	app := &freesnd{
		client: client,
		home:   home,
	}
	app.commands = map[string]commandFunc{
		"authorize":       app.authorize,
		"get-code":        app.getCode,
		"meta":            app.meta,
		"pending-uploads": app.pendingUploads,
		"refresh":         app.refreshToken,
		"upload":          app.upload,
	}
	return app, nil
}
