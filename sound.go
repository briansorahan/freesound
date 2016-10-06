package freesound

// Sound represents a sound instance.
// See https://www.freesound.org/docs/api/resources_apiv2.html#sound-instance
type Sound struct {
	ID          int      `json:"id"`
	URL         string   `json:"url"`
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	Description string   `json:"description"`
	Geotag      string   `json:"geotag"` // TODO: split into geo.Point
	Created     string   `json:"created"`
	License     string   `json:"license"`
	Type        string   `json:"type"`
	Channels    int      `json:"channels"`
	Filesize    int      `json:"filesize"` // Size of the file in bytes.
	Bitrate     int      `json:"bitrate"`
	Bitdepth    int      `json:"bitdepth"`
	Duration    int      `json:"duration"`
}
