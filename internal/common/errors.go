package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	InvalidCredentialsError = errors.New("invalid username or password")
	UserDoesNotExistError   = errors.New("user does not exist")
	UserExistError          = errors.New("user already exist")
	BindJSONError           = errors.New("error binding JSON")
	RegisterUserError       = errors.New("error register user")
	LoginUserError          = errors.New("error login user")
	ActorAlreadyExist       = errors.New("actor already exist")
	CreateActorError        = errors.New("error create actor")
)

func MakeErrorResponse(w http.ResponseWriter, status int, error error) {
	w.WriteHeader(status)

	out, _ := json.Marshal(ErrorOut{
		ErrorMessage: error.Error(),
	})

	_, _ = w.Write(out)
}
