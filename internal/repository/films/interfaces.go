package films

import "context"

type Repository interface {
	InsertNewFilm(ctx context.Context, data *CreateFilmIn) error
	ChangeFilmData(ctx context.Context, data *ChangeFilmIn) error
	DeleteFilm(ctx context.Context, data *DeleteFilmIn) error
}
