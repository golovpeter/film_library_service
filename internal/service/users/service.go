package users

import (
	"context"

	"github.com/golovpeter/vk_intership_test_task/internal/common"
	"github.com/golovpeter/vk_intership_test_task/internal/repository/users"
)

type service struct {
	repository users.Repository

	jwtKey string
}

func NewService(
	repository users.Repository,
	jwtKey string,
) UserService {
	return &service{
		repository: repository,
		jwtKey:     jwtKey,
	}
}

func (s *service) Register(ctx context.Context, data *UserDataIn) error {
	return s.repository.Register(ctx, &users.UserDataIn{
		Username: data.Username,
		Password: data.Password,
	})
}

func (s *service) Login(ctx context.Context, data *UserDataIn) (string, error) {
	userInfo, err := s.repository.GetUserInfo(ctx, &users.UserDataIn{
		Username: data.Username,
	})
	if err != nil {
		return "", err
	}

	if userInfo.Username == "" {
		return "", common.UserDoesNotExistError
	}

	correctCredentials := common.CompareHashAndPassword(data.Password, userInfo.PasswordHash)
	if !correctCredentials {
		return "", common.InvalidCredentialsError
	}

	newToken, err := common.GenerateJWT(
		s.jwtKey,
		userInfo.ID,
		userInfo.Username,
		userInfo.Role,
	)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
