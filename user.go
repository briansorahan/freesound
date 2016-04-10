package freesound

// User represents a user of the freesound API.
type User struct {
	Username string `json:"username,omitempty"`
	URL      string `json:"url,omitempty"`
	Ref      string `json:"ref,omitempty"`
}
