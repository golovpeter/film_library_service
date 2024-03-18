package films

import (
	"github.com/golovpeter/vk_intership_test_task/internal/repository/films"
	"golang.org/x/net/context"
)

type service struct {
	repository films.Repository
}

func NewService(repository films.Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateFilm(ctx context.Context, data *FilmData) (int64, error) {
	return s.repository.InsertNewFilm(ctx, &films.FilmData{
		Title:       data.Title,
		Description: data.Description,
		ReleaseDate: data.ReleaseDate,
		Rating:      data.Rating,
		Actors:      data.Actors,
	})
}

func (s *service) ChangeFilmData(ctx context.Context, data *FilmData) error {
	return s.repository.ChangeFilmData(ctx, &films.FilmData{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		ReleaseDate: data.ReleaseDate,
		Actors:      data.Actors,
	})
}

func (s *service) DeleteFilm(ctx context.Context, data *DeleteFilmIn) error {
	return s.repository.DeleteFilm(ctx, &films.DeleteFilmIn{
		FilmID: data.FilmID,
	})
}

func (s *service) GettingSortedFilms(ctx context.Context, order string) ([]*FilmData, error) {
	repoFilms, err := s.repository.GettingSortedFilms(ctx, order)
	if err != nil {
		return nil, err
	}

	filmsData := make([]*FilmData, len(repoFilms))
	for i, film := range repoFilms {
		filmsData[i] = &FilmData{
			ID:          film.ID,
			Title:       film.Title,
			Description: film.Description,
			ReleaseDate: film.ReleaseDate,
			Rating:      film.Rating,
			Actors:      film.Actors,
		}
	}

	return filmsData, nil
}

func (s *service) FindFilm(ctx context.Context, params *FindFilmIn) (*FilmData, error) {
	var repoFilm *films.FilmData
	var err error

	switch params.SearchField {
	case "title":
		repoFilm, err = s.repository.FindFilmByTitle(ctx, params.Value)
		if err != nil {
			return nil, err
		}
	case "actor":
		repoFilm, err = s.repository.FindFilmByActor(ctx, params.Value)
		if err != nil {
			return nil, err
		}
	}

	return &FilmData{
		ID:          repoFilm.ID,
		Title:       repoFilm.Title,
		Description: repoFilm.Description,
		ReleaseDate: repoFilm.ReleaseDate,
		Rating:      repoFilm.Rating,
		Actors:      repoFilm.Actors,
	}, nil
}
