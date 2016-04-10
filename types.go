package freesound

import (
	"bytes"
	"strconv"
)

// APIError represents an error that comes back from the freesound API.
type APIError struct {
	StatusCode  int    `json:"status_code,omitempty"`
	Explanation string `json:"explanation,omitempty"`
	Type        string `json:"type,omitempty"`
	Error       bool   `json:"error,omitempty"`
}

// User represents a user of the freesound API.
type User struct {
	Username string `json:"username,omitempty"`
	URL      string `json:"url,omitempty"`
	Ref      string `json:"ref,omitempty"`
}

// SoundSearchQuery represents the fields of a sound search.
type SoundSearchQuery struct {
	Query         string
	Page          int
	Filter        SoundSearchFilter
	Sort          string
	Fields        string
	SoundsPerPage int
	GroupInPacks  bool
}

// SoundSearchFilter allows you to filter the results of a sound
// search query.
type SoundSearchFilter struct {
	ID               int
	Username         string
	Created          string
	OriginalFilename string
	Description      string
	Tag              string
	License          string
	IsRemix          bool
	WasRemixed       bool
	Pack             string
	PackTokenized    string
	IsGeotagged      bool
	Type             string
	Duration         string
	Bitdepth         int
	Bitrate          int
	Samplerate       int
	Filesize         int
	Channels         int
	md5              [32]byte
	NumDownloads     int
	AvgRating        float64
	NumRatings       int
	Comment          string
	Comments         int
}

func (filt *SoundSearchFilter) String() string {
	buf := bytes.Buffer{}
	writeInt(buf, "id", filt.ID)
	writeString(buf, " username", filt.Username)
	writeString(buf, " created", filt.Created)
	writeString(buf, " original_filename", filt.OriginalFilename)
	writeString(buf, " description", filt.Description)
	writeString(buf, " tag", filt.Tag)
	writeString(buf, " license", filt.License)
	writeBool(buf, " is_remix", filt.IsRemix)
	return buf.String()
}

// writeString Write a string value to a byte buffer if it is not
// the empty string
func writeString(buf bytes.Buffer, key, val string) {
	if val != "" {
		buf.WriteString(key + ":\"" + val + "\"")
	}
}

func writeInt(buf bytes.Buffer, key string, val int) {
	if val != 0 {
		buf.WriteString(key + ":\"" + strconv.Itoa(val) + "\"")
	}
}

func writeBool(buf bytes.Buffer, key string, val bool) {
	buf.WriteString(key + ":" + strconv.FormatBool(val))
}

// SoundSearchResult represents the result of a sound search.
type SoundSearchResult struct {
	URL              string   `json:"url,omitempty"`
	OriginalFilename string   `json:"original_filename,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	Similarity       string   `json:"similarity,omitempty"`
	Serve            string   `json:"serve,omitempty"`
	Type             string   `json:"type,omitempty"`
	Ref              string   `json:"ref,omitempty"`
	ID               int      `json:"id,omitempty"`
	Pack             string   `json:"pack,omitempty"`
}
