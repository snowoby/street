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

func WTF(err error) HTTPError {
	return HTTPError{
		code:    http.StatusBadRequest,
		message: err.Error(),
	}
}

func DatabaseError(err error) HTTPError {
	code := http.StatusBadGateway
	switch err.(type) {
	case *ent.NotFoundError:
		code = http.StatusNotFound
	case *ent.NotLoadedError:
		code = http.StatusBadGateway
	case *ent.NotSingularError:
		code = http.StatusBadGateway
	case *ent.ValidationError:
		code = http.StatusBadRequest
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
	switch err.(type) {
	case HTTPError:
		e := err.(HTTPError)
		return e
	case *ent.NotFoundError:
	case *ent.NotLoadedError:
	case *ent.NotSingularError:
	case *ent.ValidationError:
	case *ent.ConstraintError:
		return DatabaseError(err)
	default:
		return WTF(err)
	}
	return HTTPError{}
}

var (
	NotFoundError = HTTPError{code: http.StatusNotFound, message: "not found"}
)
