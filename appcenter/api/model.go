package api

// ErrorResponse struct representing API error
type ErrorResponse struct {
	Error Error `json:"error,omitEmpty"`
}

// Error struct representing API error
type Error struct {
	Code    string `json:"code,omitEmpty"`
	Message string `json:"message,omitEmpty"`
}
