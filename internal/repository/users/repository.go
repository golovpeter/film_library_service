package users

import (
	"context"

	"github.com/golovpeter/vk_intership_test_task/internal/common"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) Repository {
	return &repository{conn: conn}
}

const registerUserQuery = `
	INSERT INTO users(username, password_hash)
	VALUES ($1, $2)
	ON CONFLICT DO NOTHING 
`

func (r *repository) Register(ctx context.Context, data *UserDataIn) error {
	res, err := r.conn.ExecContext(ctx, registerUserQuery, data.Username, data.Password)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return common.UserExistError
	}

	return nil
}

const getUserQuery = `
	SELECT id, username, password_hash, role
	FROM users
	WHERE username = $1
`

func (r *repository) GetUserInfo(ctx context.Context, data *UserDataIn) (*UserDataOut, error) {
	var out UserDataOut

	err := r.conn.GetContext(ctx, &out, getUserQuery, data.Username)
	if err != nil {
		return nil, err
	}

	return &out, nil
}
