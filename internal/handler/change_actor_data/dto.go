package change_actor_data

type ChangeActorDataIn struct {
	ID        int64  `json:"id" example:"10"`
	Name      string `json:"name,omitempty" example:"Cillian Murphy"`
	Gender    string `json:"gender,omitempty" example:"male"`
	BirthDate string `json:"birth_date,omitempty" example:"1976-05-25"`
}
