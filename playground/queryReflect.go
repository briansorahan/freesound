package main

import (
	"fmt"
	"reflect"
)

type SoundSearchFilter struct {
	Id               int      `query:"id,omitempty"`
	Username         string   `query:"username,omitempty"`
	Created          string   `query:"created,omitempty"`
	OriginalFilename string   `query:"original_filename,omitempty"`
	Description      string   `query:"description,omitempty"`
	Tag              string   `query:"tag,omitempty"`
	License          string   `query:"license,omitempty"`
	IsRemix          bool     `query:"is_remix,omitempty"`
	WasRemixed       bool     `query:"was_remixed,omitempty"`
	Pack             string   `query:"pack,omitempty"`
	PackTokenized    string   `query:"pack_tokenized,omitempty"`
	IsGeotagged      bool     `query:"is_geotagged,omitempty"`
	Type             string   `query:"type,omitempty"`
	Duration         string   `query:"duration,omitempty"`
	Bitdepth         int      `query:"bitdepth,omitempty"`
	Bitrate          int      `query:"bitrate,omitempty"`
	Samplerate       int      `query:"samplerate,omitempty"`
	Filesize         int      `query:"filesize,omitempty"`
	Channels         int      `query:"channels,omitempty"`
	md5              [32]byte `query:"md5,omitempty"`
	NumDownloads     int      `query:"num_downloads,omitempty"`
	AvgRating        float64  `query:"avg_rating,omitempty"`
	NumRatings       int      `query:"num_ratings,omitempty"`
	Comment          string   `query:"comment,omitempty"`
	Comments         int      `query:"comments,omitempty"`
}

func main() {
	filt := SoundSearchFilter{}
	t := reflect.TypeOf(filt)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag
	}
}
