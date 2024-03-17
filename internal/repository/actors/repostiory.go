package actors

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

const insertActorQuery = `
	INSERT INTO actors(name, gender, birth_date) 
	VALUES ($1 , $2, $3)
`

func (r *repository) CreateActor(ctx context.Context, data *ActorDataIn) error {
	res, err := r.conn.ExecContext(ctx, insertActorQuery, data.Name, data.Gender, data.BirthDate)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return common.ActorAlreadyExist
	}

	return nil
}

func (r *repository) ChangeActorInfo(ctx context.Context, data *ActorDataIn) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) DeleteActor(ctx context.Context, data *DeleteActorIn) error {
	//TODO implement me
	panic("implement me")
}
