package create_actor

type CreateActorIn struct {
	Name      string `json:"name" example:"Cillian Murphy"`
	Gender    string `json:"gender" example:"male"`
	BirthDate string `json:"birth_date" example:"1976-05-25"`
}

type CreateActorOut struct {
	ID int64 `json:"id" example:"10"`
}
