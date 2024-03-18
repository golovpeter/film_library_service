package films

import "golang.org/x/net/context"

type FilmsService interface {
	CreateFilm(ctx context.Context, data *CreateFilmIn) error
	ChangeFilmData(ctx context.Context, data *ChangeFilmIn) error
	DeleteFilm(ctx context.Context, data *DeleteFilmIn) error
}
