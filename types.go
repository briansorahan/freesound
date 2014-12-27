package freesound

type User struct {
	Username string `json:"username,omitempty"`
	URL      string `json:"url,omitempty"`
	Ref      string `json:"ref,omitempty"`
}

type SoundSearchResult struct {
	URL              string   `json:"url,omitempty"`
	OriginalFilename string   `json:"original_filename,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	Similarity       string   `json:"similarity,omitempty"`
	Serve            string   `json:"serve,omitempty"`
	Type             string   `json:"type,omitempty"`
	Ref              string   `json:"ref,omitempty"`
	Id               int      `json:"id,omitempty"`
	Pack             string   `json:"pack,omitempty"`
}

type Client interface {
	SoundSearch(query string) (*SoundSearchResult, error)
	Version() int
}
