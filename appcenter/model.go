package appcenter

// ErrorResponse struct representing API error
type ErrorResponse struct {
	Error Error `json:"error,omitEmpty"`
}

// Error struct representing API error
type Error struct {
	code    string `json:"code,omitEmpty"`
	message string `json:"message,omitEmpty"`
}
