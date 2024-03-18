package films

import "golang.org/x/net/context"

type FilmsService interface {
	CreateFilm(ctx context.Context, data *FilmData) (int64, error)
	ChangeFilmData(ctx context.Context, data *FilmData) error
	DeleteFilm(ctx context.Context, data *DeleteFilmIn) error
	GettingSortedFilms(ctx context.Context, order string) ([]*FilmData, error)
	FindFilm(ctx context.Context, params *FindFilmIn) (*FilmData, error)
}
