package web

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Error is used to pass an error during the request through the
// application with web specific context.
type Error struct {
	Err    error
	Status int
}

// implementing the error interface
func (err *Error) Error() string {
	return err.Err.Error()
}

//NewRequestError is our trusted error // this func should be used to pass errors back to the client (postman/frontend)
func NewRequestError(err error, status int) error {
	return &Error{Err: err, Status: status}
}
