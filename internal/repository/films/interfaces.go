package films

import "context"

type Repository interface {
	CreateFilm(ctx context.Context, data *CreateFilmIn) error
	ChangeFilmData() error
	DeleteFilm() error
}
