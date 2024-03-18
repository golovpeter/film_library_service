package films

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

const insertNewFilmQuery = `
	INSERT INTO films(title, description, release_date, rating)
	VALUES ($1, $2, $3, $4)
	RETURNING id
`

const insertNewFilmActorsQuery = `
	INSERT INTO films_and_actors(film_id, actor_id) 
	VALUES %s
`

func (r *repository) InsertNewFilm(ctx context.Context, data *FilmData) error {
	actorsIDs, err := r.getActorsIDs(ctx, data.Actors)

	if len(data.Actors) != len(actorsIDs) {
		return common.UnknownActorError
	}

	tx, err := r.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
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

	insertActorsQuery, valueArgs := r.getInsertActorsQuery(actorsIDs, newFilmID)

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

const deleteCurActorsQuery = `
	DELETE FROM films_and_actors
	WHERE film_id = $1
`

func (r *repository) ChangeFilmData(ctx context.Context, data *FilmData) error {
	var queryBuilder strings.Builder
	var args []interface{}
	argIdx := 1

	queryBuilder.WriteString("UPDATE films SET ")

	if data.Title != "" {
		queryBuilder.WriteString(fmt.Sprintf("title = $%d, ", argIdx))
		args = append(args, data.Title)
		argIdx++
	}

	if data.Description != "" {
		queryBuilder.WriteString(fmt.Sprintf("description = $%d, ", argIdx))
		args = append(args, data.Description)
		argIdx++
	}

	if data.ReleaseDate != "" {
		queryBuilder.WriteString(fmt.Sprintf("release_date = $%d, ", argIdx))
		args = append(args, data.ReleaseDate)
		argIdx++
	}

	if data.Rating != -1 {
		queryBuilder.WriteString(fmt.Sprintf("rating = $%d, ", argIdx))
		args = append(args, data.Rating)
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

	if len(data.Actors) == 0 {
		return nil
	}

	actorsIDs, err := r.getActorsIDs(ctx, data.Actors)
	if err != nil {
		return err
	}

	if len(data.Actors) != len(actorsIDs) {
		return common.UnknownActorError
	}

	tx, err := r.conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, deleteCurActorsQuery, data.ID)
	if err != nil {
		return err
	}

	insertActorsQuery, valueArgs := r.getInsertActorsQuery(actorsIDs, data.ID)

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

const deleteFilmQuery = `
	DELETE FROM films
	WHERE id = $1
`

func (r *repository) DeleteFilm(ctx context.Context, data *DeleteFilmIn) error {
	res, err := r.conn.ExecContext(ctx, deleteFilmQuery, data.FilmID)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return common.FilmDoesNotExistError
	}

	return nil
}

func (r *repository) GettingSortedFilms(ctx context.Context, order string) ([]*FilmData, error) {
	var gettingFilmsQuery = `
	SELECT id, title, description, release_date, rating
	FROM films
	ORDER BY 
`

	if order == "" {
		gettingFilmsQuery += "rating"
	} else {
		gettingFilmsQuery += order
	}

	if order != "title" {
		gettingFilmsQuery += " desc"
	}

	var result []*FilmData

	err := r.conn.SelectContext(ctx, &result, gettingFilmsQuery)
	if err != nil {
		return nil, err
	}

	for _, res := range result {
		actors, err := r.getFilmActors(ctx, res.ID)
		if err != nil {
			return nil, err
		}

		res.Actors = actors
	}

	return result, nil
}

const findFilmByTitle = `
	SELECT id, title, description, release_date, rating
	FROM films
	WHERE LOWER(title) LIKE LOWER('%' || $1 || '%')
	LIMIT 1
`

func (r *repository) FindFilmByTitle(ctx context.Context, title string) (*FilmData, error) {
	var result []*FilmData

	err := r.conn.SelectContext(ctx, &result, findFilmByTitle, title)
	if err != nil {
		return nil, err
	}

	if len(result) != 0 {
		return result[0], nil
	} else {
		return nil, common.NoMatchesFoundError
	}
}

const findFilmByActor = `
	SELECT films.title, films.description, films.release_date, films.rating
	FROM FILMS
	JOIN films_and_actors ON films.id = films_and_actors.film_id
	JOIN actors ON films_and_actors.actor_id = actors.id
	WHERE lower(actors.name) LIKE lower('%' || $1 || '%')
	LIMIT 1;
`

func (r *repository) FindFilmByActor(ctx context.Context, actor string) (*FilmData, error) {
	var result []*FilmData

	err := r.conn.SelectContext(ctx, &result, findFilmByActor, actor)
	if err != nil {
		return nil, err
	}

	if len(result) != 0 {
		return result[0], nil
	} else {
		return nil, common.NoMatchesFoundError
	}
}

const getActorsIDQuery = `
	SELECT id
	FROM actors
	WHERE name in (?)
`

func (r *repository) getActorsIDs(ctx context.Context, actors []string) ([]int64, error) {
	var actorsIDs []int64

	query, args, err := sqlx.In(getActorsIDQuery, actors)
	if err != nil {
		return nil, err
	}

	query = r.conn.Rebind(query)

	err = r.conn.SelectContext(ctx, &actorsIDs, query, args...)
	if err != nil {
		return nil, err
	}

	return actorsIDs, nil
}

func (r *repository) getInsertActorsQuery(actorsIDs []int64, filmID int64) (string, []interface{}) {
	var valueStrings []string
	var valueArgs []interface{}

	for _, actorID := range actorsIDs {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, filmID, actorID)
	}

	query := fmt.Sprintf(insertNewFilmActorsQuery, strings.Join(valueStrings, ","))
	query = r.conn.Rebind(query)

	return query, valueArgs
}

const getFilmActorsQuery = `
	SELECT actors.name
	FROM films
         JOIN
     films_and_actors ON films.id = films_and_actors.film_id
         JOIN
     actors ON films_and_actors.actor_id = actors.id
	WHERE films.id = $1
`

func (r *repository) getFilmActors(ctx context.Context, filmID int64) ([]string, error) {
	var actors []string

	err := r.conn.SelectContext(ctx, &actors, getFilmActorsQuery, strconv.FormatInt(filmID, 10))
	if err != nil {
		return nil, err
	}

	return actors, nil
}
