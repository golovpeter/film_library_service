package films

type FilmData struct {
	ID          int64
	Title       string `db:"title"`
	Description string `db:"description"`
	ReleaseDate string `db:"release_date"`
	Rating      int    `db:"rating"`
	Actors      []string
}

type DeleteFilmIn struct {
	FilmID int64
}
