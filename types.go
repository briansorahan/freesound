package freesound

import (
	"bytes"
	"strconv"
)

// writeString writes a string value to a byte buffer if it is not the empty string.
func writeString(buf bytes.Buffer, key, val string) {
	if val != "" {
		buf.WriteString(key + ":\"" + val + "\"")
	}
}

// writeInt writes a int value to a byte buffer if it is not 0.
func writeInt(buf bytes.Buffer, key string, val int) {
	if val != 0 {
		buf.WriteString(key + ":\"" + strconv.Itoa(val) + "\"")
	}
}

// writeBool writes a bool value to a byte buffer.
func writeBool(buf bytes.Buffer, key string, val bool) {
	buf.WriteString(key + ":" + strconv.FormatBool(val))
}
