package actors

type ActorData struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	Gender    string `db:"gender"`
	BirthDate string `db:"birth_date"`
	Films     []string
}

type ChangeActorDataIn struct {
	ID        int64
	Name      string
	Gender    string
	BirthDate string
}

type DeleteActorIn struct {
	ActorID int64
}
