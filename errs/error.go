package errs

import (
	"net/http"
	"street/ent"
)

type HTTPError struct {
	code    int    `json:"code"`
	message string `json:"message"`
}

type ResponseError interface {
	Code() int
	Error() string
	Message() string
}

func (e HTTPError) Code() int {
	return e.code
}

func (e HTTPError) Message() string {
	return e.message
}

func (e HTTPError) Error() string {
	return e.message
}

func BindingError(err error) HTTPError {
	return HTTPError{
		code:    http.StatusBadRequest,
		message: err.Error(),
	}
}

func WTF(info string) HTTPError {
	return HTTPError{
		code:    http.StatusBadRequest,
		message: info,
	}
}

func DatabaseError(err error) HTTPError {
	code := http.StatusBadGateway
	switch err.(type) {
	case *ent.NotFoundError:
		code = http.StatusNotFound
	case *ent.ConstraintError:
		code = http.StatusBadRequest
	}

	return HTTPError{
		code:    code,
		message: err.Error(),
	}
}

func Detect(err error) HTTPError {
	// TODO
	return HTTPError{}
}

var (
	NotFoundError = HTTPError{code: http.StatusNotFound, message: "not found"}
)
