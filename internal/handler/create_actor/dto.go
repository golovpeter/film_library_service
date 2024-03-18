package create_actor

type CreateActorIn struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
}

type CreateActorOut struct {
	ID int64 `json:"id"`
}
