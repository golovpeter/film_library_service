package films

type CreateFilmIn struct {
	Title       string
	Description string
	ReleaseDate string
	Rating      int
	Actors      []string
}
