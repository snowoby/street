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

func WTF(err error) HTTPError {
	return HTTPError{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
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
		Code:    code,
		Message: err.Error(),
	}
}

func Detect(err error) HTTPError {
	// TODO
	switch t := err.(type) {
	case HTTPError:
		return t
	case *ent.NotFoundError, *ent.NotLoadedError, *ent.NotSingularError, *ent.ValidationError, *ent.ConstraintError:
		return DatabaseError(t)
	default:
		return WTF(err)
	}
}

var (
	NotFoundError = HTTPError{Code: http.StatusNotFound, Message: "not found"}
)
