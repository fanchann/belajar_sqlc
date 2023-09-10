package exception

type HttpError struct {
	message    string
	statusCode uint
}

func (h *HttpError) Error() string {
	return h.message
}

func (h *HttpError) StatusCode() uint {
	return h.statusCode
}

func NewHTTPError(statusCode uint, message string) *HttpError {
	return &HttpError{message, statusCode}
}

type ValidationError struct {
	Err error
}

func (e *ValidationError) Error() string {
	return e.Err.Error()
}
func NewValidationError(err error) *ValidationError {
	return &ValidationError{Err: err}
}

type NotFoundErr struct {
	Err error
}

func (n *NotFoundErr) Error() string {
	return n.Err.Error()
}

func NewNotFoundErr(err error) *NotFoundErr {
	return &NotFoundErr{Err: err}
}
