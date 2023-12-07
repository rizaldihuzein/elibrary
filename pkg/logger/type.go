package logger

type (
	Log struct {
		Type    string `json:"type"`
		Error   string `json:"error,omitempty"`
		Message string `json:"message"`
	}
)
