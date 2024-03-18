package actors

import (
	"context"
	"fmt"
	"strconv"
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
	RETURNING id
`

func (r *repository) CreateActor(ctx context.Context, data *ActorData) (int64, error) {
	var newActorId int64

	err := r.conn.QueryRowContext(ctx,
		insertActorQuery,
		data.Name,
		data.Gender,
		data.BirthDate,
	).Scan(&newActorId)
	if err != nil {
		return 0, err
	}

	return newActorId, err
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

const deleteActorQuery = `
	DELETE FROM actors
	WHERE id = $1
`

func (r *repository) DeleteActor(ctx context.Context, data *DeleteActorIn) error {
	res, err := r.conn.ExecContext(ctx, deleteActorQuery, data.ActorID)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return common.ActorDoesNotExistError
	}

	return nil
}

const getAllActorsQuery = `
	SELECT id, name, gender, birth_date
	FROM actors
`

func (r *repository) GetAllActors(ctx context.Context) ([]*ActorData, error) {
	var actors []*ActorData

	err := r.conn.SelectContext(ctx, &actors, getAllActorsQuery)
	if err != nil {
		return nil, err
	}

	for _, actor := range actors {
		actorFilms, err := r.getActorFilms(ctx, actor.ID)
		if err != nil {
			return nil, err
		}

		actor.Films = actorFilms
	}

	return actors, nil
}

const getActorFilmsQuery = `
	SELECT title
	FROM films_and_actors
         JOIN films ON films_and_actors.film_id = films.id
	WHERE actor_id = $1;
`

func (r *repository) getActorFilms(ctx context.Context, actorID int64) ([]string, error) {
	var films []string

	err := r.conn.SelectContext(ctx, &films, getActorFilmsQuery, strconv.FormatInt(actorID, 10))
	if err != nil {
		return nil, err
	}

	return films, nil
}
