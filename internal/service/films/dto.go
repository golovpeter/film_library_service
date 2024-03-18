package films

type FilmData struct {
	ID          int64
	Title       string
	Description string
	ReleaseDate string
	Rating      int
	Actors      []string
}

type DeleteFilmIn struct {
	FilmID int64
}

type FindFilmIn struct {
	SearchField string
	Value       string
}
