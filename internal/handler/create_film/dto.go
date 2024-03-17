package create_film

type CreatFilmIn struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ReleaseDate string   `json:"release_date"`
	Rating      int      `json:"rating"`
	Actors      []string `json:"actors"`
}
