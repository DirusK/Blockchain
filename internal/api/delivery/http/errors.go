package http

type (
	// ErrResponse response with error.
	ErrResponse struct {
		Error string `json:"error"`
	}
)
