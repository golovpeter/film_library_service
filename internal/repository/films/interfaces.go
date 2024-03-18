package films

import "context"

type Repository interface {
	InsertNewFilm(ctx context.Context, data *FilmData) error
	ChangeFilmData(ctx context.Context, data *FilmData) error
	DeleteFilm(ctx context.Context, data *DeleteFilmIn) error
	GettingSortedFilms(ctx context.Context, order string) ([]*FilmData, error)
}
