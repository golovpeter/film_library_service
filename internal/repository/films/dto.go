package films

type CreateFilmIn struct {
	Title       string
	Description string
	ReleaseDate string
	Rating      int
	Actors      []string
}

type ChangeFilmIn struct {
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
