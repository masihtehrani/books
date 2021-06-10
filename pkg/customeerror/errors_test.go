package customeerror_test

import (
	"testing"

	"github.com/masihtehrani/books/pkg/customeerror"
)

func TestError_Error(t *testing.T) {
	err := customeerror.Error{Err: "sample error message", Code: "10"}
	msg := "sample error message"
	// New := customeerror.Error(msg)

	if msg != err.Error() {
		t.Error("msg not equal message error")
	}
}
