package response

// ErrorResponse represents the JSON response for an error
type ErrorResponse struct {
	Status  int `json:"status"`
	Message string `json:"message"`
}
