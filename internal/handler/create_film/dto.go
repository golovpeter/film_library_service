package create_film

type CreatFilmIn struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ReleaseDate string   `json:"release_date"`
	Rating      int      `json:"rating"`
	Actors      []string `json:"actors"`
}

type CreateFilmOut struct {
	ID int64 `json:"id"`
}
