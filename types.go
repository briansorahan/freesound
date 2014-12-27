package freesound

import (
	"bytes"
	"strconv"
)

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
	Filter        SoundSearchFilter
	Sort          string
	Fields        string
	SoundsPerPage int
	GroupInPacks  bool
}

type SoundSearchFilter struct {
	Id               int
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

func (self *SoundSearchFilter) String() string {
	buf := bytes.Buffer{}
	writeInt(buf,    "id", self.Id)
	writeString(buf, " username", self.Username)
	writeString(buf, " created", self.Created)
	writeString(buf, " original_filename", self.OriginalFilename)
	writeString(buf, " description", self.Description)
	writeString(buf, " tag", self.Tag)
	writeString(buf, " license", self.License)
	writeBool(buf,   " is_remix", self.IsRemix)
	return buf.String()
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
