package get_all_actors

type ActorData struct {
	ID        int64    `json:"id" example:"10"`
	Name      string   `json:"name" example:"Leonardo Dicaprio"`
	Gender    string   `json:"gender" example:"male"`
	BirthDate string   `json:"birth_date" example:"1995-03-23"`
	Films     []string `example:"Titanic,The Revenant"`
}
