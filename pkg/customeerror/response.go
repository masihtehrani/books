package customeerror

import (
	"errors"
)

type ResErr struct {
	Code   string `json:"code"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

type ResponseError struct {
	Errors []ResErr `json:"errors"`
}

func NewErrorResponse(err error, status int) *ResponseError {
	for {
		e := errors.Unwrap(err)
		if e == nil {
			break
		}

		err = e
	}

	res := make([]ResErr, 0)

	switch err := err.(type) { //nolint: errorlint
	case Error:
		res = append(res, ResErr{
			Code:   err.Code,
			Detail: err.Error(),
			Status: status,
		})

		return &ResponseError{Errors: res}

	default:
		res = append(res, ResErr{
			Status: status,
			Detail: err.Error(),
		})

		return &ResponseError{Errors: res}
	}
}
