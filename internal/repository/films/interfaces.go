package films

import "context"

type Repository interface {
	InsertNewFilm(ctx context.Context, data *FilmData) (int64, error)
	ChangeFilmData(ctx context.Context, data *FilmData) error
	DeleteFilm(ctx context.Context, data *DeleteFilmIn) error
	GettingSortedFilms(ctx context.Context, order string) ([]*FilmData, error)
	FindFilmByTitle(ctx context.Context, title string) (*FilmData, error)
	FindFilmByActor(ctx context.Context, actor string) (*FilmData, error)
}
