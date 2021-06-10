package customeerror

import (
	"fmt"
)

type Error struct {
	Err  string
	Code string
}

func (e Error) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func New(code string, s string) Error {
	return Error{
		Err:  s,
		Code: code,
	}
}
