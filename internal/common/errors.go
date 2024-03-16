package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	UserExistError    = errors.New("user already exist")
	BindJSONError     = errors.New("error binding JSON")
	RegisterUserError = errors.New("error register user")
)

func MakeErrorResponse(w http.ResponseWriter, status int, error error) {
	w.WriteHeader(status)

	out, _ := json.Marshal(ErrorOut{
		ErrorMessage: error.Error(),
	})

	_, _ = w.Write(out)
}
