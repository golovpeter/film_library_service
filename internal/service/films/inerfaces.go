package films

import "golang.org/x/net/context"

type FilmsService interface {
	CreateFilm(ctx context.Context, data *CreateFilmIn) error
	ChangeFilmData() error
	DeleteFilm() error
}
