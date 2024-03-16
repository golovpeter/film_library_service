package users

import (
	"context"

	"github.com/golovpeter/vk_intership_test_task/internal/repository/users"
)

type service struct {
	repository users.Repository
}

func NewService(repository users.Repository) UserService {
	return &service{repository: repository}
}

func (s *service) Register(ctx context.Context, data *UserDataIn) error {
	return s.repository.Register(ctx, &users.UserDataIn{
		Username: data.Username,
		Password: data.Password,
	})
}

func (s *service) Login(ctx context.Context, data *UserDataIn) (string, error) {
	//TODO implement me
	panic("implement me")
}
