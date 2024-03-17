package change_film_data

type ChangeFilmIn struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	ReleaseDate string   `json:"release_date,omitempty"`
	Rating      int      `json:"rating,omitempty"`
	Actors      []string `json:"actors,omitempty"`
}
