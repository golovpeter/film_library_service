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
	DeleteActorError        = errors.New("error delete actor")
	UnknownActorError       = errors.New("adding an unknown actor not possible")
	CreateFilmError         = errors.New("create film error")
	ChangeFilmDataError     = errors.New("change film data error")
	DeleteFilmError         = errors.New("delete film error")
	FilmDoesNotExistError   = errors.New("film doest not exist")
	GettingFilmsError       = errors.New("error get films")
)

func MakeErrorResponse(w http.ResponseWriter, status int, error error) {
	w.WriteHeader(status)

	out, _ := json.Marshal(ErrorOut{
		ErrorMessage: error.Error(),
	})

	_, _ = w.Write(out)
}
