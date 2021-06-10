package structs

import "github.com/masihtehrani/books/pkg/customeerror"

// nolint: revive
var (
	ErrUnknown          = customeerror.New("1001", "Unknown error")
	ErrEmptyIPANDPORT   = customeerror.New("1002", "some fields are empty in cli: IP, Port, routers")
	ErrJwtSecretToken   = customeerror.New("1003", "secret token not be filled")
	ErrUserID           = customeerror.New("1004", "user id is not filled")
	ErrValidationSignUp = customeerror.New("1005", "in sign up must filled all fileds username,password"+
		",full_name & pseudonym")
	ErrValidationSigIn      = customeerror.New("1006", "in sign in must filled all fileds username & password")
	ErrValidationCreateBook = customeerror.New("1007", "in create book must filled all fileds title & description")
	ErrUnauthorized         = customeerror.New("1008", "Unauthorized")
)
