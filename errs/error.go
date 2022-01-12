package errs

import (
	"net/http"
	"street/ent"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseError interface {
	Code() int
	Error() string
	Message() string
}

func (e HTTPError) Error() string {
	return e.Message
}

func BindingError(err error) HTTPError {
	return HTTPError{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
}

func WTF(info string) HTTPError {
	return HTTPError{
		Code:    http.StatusBadRequest,
		Message: info,
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
		Code:    code,
		Message: err.Error(),
	}
}

var (
	NotFoundError = HTTPError{Code: http.StatusNotFound, Message: "not found"}
)
