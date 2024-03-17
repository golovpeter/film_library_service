package films

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

const getActorsIDQuery = `
	SELECT id
	FROM actors
	WHERE name in (?)
`

const insertNewFilmQuery = `
	INSERT INTO films(title, description, release_date, rating)
	VALUES ($1, $2, $3, $4)
	RETURNING id
`

const insertNewFilmActorsQuery = `
	INSERT INTO films_and_actors(film_id, actor_id) 
	VALUES %s
`

func (r *repository) CreateFilm(ctx context.Context, data *CreateFilmIn) error {
	var actorsIDs []int64

	query, args, err := sqlx.In(getActorsIDQuery, data.Actors)
	if err != nil {
		return err
	}

	query = r.conn.Rebind(query)

	tx, err := r.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var actorID int64

		if err = rows.Scan(&actorID); err != nil {
			_ = tx.Rollback()
			return err
		}

		actorsIDs = append(actorsIDs, actorID)
	}

	if len(data.Actors) != len(actorsIDs) {
		return common.UnknownActorError
	}

	var newFilmID int64
	err = tx.QueryRowContext(ctx, insertNewFilmQuery,
		data.Title,
		data.Description,
		data.ReleaseDate,
		data.Rating,
	).Scan(&newFilmID)
	if err != nil {
		return err
	}

	var valueStrings []string
	var valueArgs []interface{}

	for _, actorID := range actorsIDs {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, newFilmID, actorID)
	}

	insertActorsQuery := fmt.Sprintf(insertNewFilmActorsQuery, strings.Join(valueStrings, ","))
	insertActorsQuery = r.conn.Rebind(insertActorsQuery)

	_, err = tx.ExecContext(ctx, insertActorsQuery, valueArgs...)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) ChangeFilmData() error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) DeleteFilm() error {
	//TODO implement me
	panic("implement me")
}
