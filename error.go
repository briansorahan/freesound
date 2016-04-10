package freesound

// APIError represents an error that comes back from the freesound API.
type APIError struct {
	StatusCode  int    `json:"status_code,omitempty"`
	Explanation string `json:"explanation,omitempty"`
	Type        string `json:"type,omitempty"`
	Error       bool   `json:"error,omitempty"`
}
