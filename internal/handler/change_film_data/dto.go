package change_film_data

type ChangeFilmIn struct {
	ID          int64    `json:"id" example:"3"`
	Title       string   `json:"title,omitempty" example:"Titanic"`
	Description string   `json:"description,omitempty" example:"Description"`
	ReleaseDate string   `json:"release_date,omitempty" example:"1998-02-20"`
	Rating      int      `json:"rating,omitempty" example:"10"`
	Actors      []string `json:"actors,omitempty" example:"Leonardo Dicaprio"`
}
