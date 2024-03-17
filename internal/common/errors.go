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
	CreateActorError        = errors.New("error create actor")
	ChangeActorDataError    = errors.New("error change actor data")
	ActorDoesNotExistError  = errors.New("actor does not exist")
)

func MakeErrorResponse(w http.ResponseWriter, status int, error error) {
	w.WriteHeader(status)

	out, _ := json.Marshal(ErrorOut{
		ErrorMessage: error.Error(),
	})

	_, _ = w.Write(out)
}
