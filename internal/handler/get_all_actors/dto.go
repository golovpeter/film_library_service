package get_all_actors

type ActorData struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Films     []string
}
