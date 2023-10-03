package model

type Response struct {
	Data          interface{}
	ErrorResponse ErrorResponse
}

type ErrorResponse struct {
	Err         error
	Description string
	StatusCode  int
}
