package errors

import (
	"net/http"
	"street/ent"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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
	}

	return HTTPError{
		Code:    code,
		Message: err.Error(),
	}
}
