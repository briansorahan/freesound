package freesound

type ApiError struct {
	StatusCode  int    `json:"status_code,omitempty"`
	Explanation string `json:"explanation,omitempty"`
	Type        string `json:"type,omitempty"`
	Error       bool   `json:"error,omitempty"`
}

type User struct {
	Username string `json:"username,omitempty"`
	URL      string `json:"url,omitempty"`
	Ref      string `json:"ref,omitempty"`
}

type SoundSearchQuery struct {
	Query         string
	Page          int
	Filter        string
	Sort          string
	Fields        string
	SoundsPerPage int
	GroupInPacks  bool
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
	SoundSearch(query SoundSearchQuery) (*SoundSearchResult, error)
	Version() int
}
