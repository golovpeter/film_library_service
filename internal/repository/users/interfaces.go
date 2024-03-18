package users

import "context"

type Repository interface {
	Register(ctx context.Context, data *UserDataIn) error
	GetUserInfo(ctx context.Context, data *UserDataIn) (*UserDataOut, error)
	GetUserRole(ctx context.Context, username string) (string, error)
}
