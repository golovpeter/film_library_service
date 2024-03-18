package films

type FilmData struct {
	ID          int64    `db:"id"`
	Title       string   `db:"title"`
	Description string   `db:"description"`
	ReleaseDate string   `db:"release_date"`
	Rating      int      `db:"rating"`
	Actors      []string `db:"-"`
}

type DeleteFilmIn struct {
	FilmID int64
}

type FindFilmIn struct {
	Title      string
	ActorName  string
	SearchType string
}
