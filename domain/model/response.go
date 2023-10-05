package model

type Response struct {
	Data          interface{}   `json:"data,omitempty"`
	ErrorResponse ErrorResponse `json:"error_response,omitempty"`
}

type ErrorResponse struct {
	Err         error  `json:"-"`
	Description string `json:"description,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
}
