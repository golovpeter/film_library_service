package get_sorted_films

type FilmData struct {
	ID          int64    `json:"id" example:"10"`
	Title       string   `json:"title" example:"Titanic"`
	Description string   `json:"description" example:"Description"`
	ReleaseDate string   `json:"release_date" example:"1998-02-20"`
	Rating      int      `json:"rating" example:"10"`
	Actors      []string `json:"actors" example:"Leonardo Dicaprio"`
}
