package users

import "context"

type Repository interface {
	Register(ctx context.Context, data *UserDataIn) error
	Login(ctx context.Context, data *UserDataIn) (string, error)
}
