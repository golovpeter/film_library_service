package change_actor_data

type ChangeActorDataIn struct {
	ID        int64  `json:"id"`
	Name      string `json:"name,omitempty"`
	Gender    string `json:"gender,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
}
