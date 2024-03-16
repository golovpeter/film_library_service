package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	UserExistError          = errors.New("user already exist")
	UserDoesNotExistError   = errors.New("user does not exist")
	InvalidCredentialsError = errors.New("invalid username or password")
	BindJSONError           = errors.New("error binding JSON")
	RegisterUserError       = errors.New("error register user")
	LoginUserError          = errors.New("error login user")
)

func MakeErrorResponse(w http.ResponseWriter, status int, error error) {
	w.WriteHeader(status)

	out, _ := json.Marshal(ErrorOut{
		ErrorMessage: error.Error(),
	})

	_, _ = w.Write(out)
}
