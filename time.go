package freesound

import (
	"strings"
	"time"
)

const timeLayout = "2006-01-02T15:04:05"

// Time is a utility type for marshalling and unmarshalling timestamps.
// Timestamps in the freesound API don't have a 'Z' at then end, so
// we can not use the std library's time.Time.
type Time time.Time

// MarshalJSON marshals a Time to JSON.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format(timeLayout) + `"`), nil
}

// UnmarshalJSON unmarshals a Time from JSON.
func (t *Time) UnmarshalJSON(bs []byte) error {
	s := strings.Trim(string(bs), `"`)
	tt, err := time.Parse(timeLayout, s)
	if err != nil {
		return err
	}
	*t = Time(tt)
	return nil
}
