package actors

import (
	"context"
	"fmt"
	"strings"

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
	_, err := r.conn.ExecContext(ctx, insertActorQuery, data.Name, data.Gender, data.BirthDate)
	return err
}

func (r *repository) ChangeActorInfo(ctx context.Context, data *ChangeActorDataIn) error {
	var queryBuilder strings.Builder
	var args []interface{}
	argIdx := 1

	queryBuilder.WriteString("UPDATE actors SET ")

	if data.Name != "" {
		queryBuilder.WriteString(fmt.Sprintf("name = $%d, ", argIdx))
		args = append(args, data.Name)
		argIdx++
	}

	if data.Gender != "" {
		queryBuilder.WriteString(fmt.Sprintf("gender = $%d, ", argIdx))
		args = append(args, data.Gender)
		argIdx++
	}

	if data.BirthDate != "" {
		queryBuilder.WriteString(fmt.Sprintf("birth_date = $%d, ", argIdx))
		args = append(args, data.BirthDate)
		argIdx++
	}

	query := strings.TrimSuffix(queryBuilder.String(), ", ")

	if len(args) > 0 {
		query += fmt.Sprintf(" WHERE id = $%d", argIdx)
		args = append(args, data.ID)

		res, err := r.conn.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return common.ActorDoesNotExistError
		}
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
