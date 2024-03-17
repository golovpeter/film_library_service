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

func (s *service) CreateFilm(ctx context.Context, data *CreateFilmIn) error {
	return s.repository.CreateFilm(ctx, &films.CreateFilmIn{
		Title:       data.Title,
		Description: data.Description,
		ReleaseDate: data.ReleaseDate,
		Rating:      data.Rating,
		Actors:      data.Actors,
	})
}

func (s *service) ChangeFilmData() error {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteFilm() error {
	//TODO implement me
	panic("implement me")
}
