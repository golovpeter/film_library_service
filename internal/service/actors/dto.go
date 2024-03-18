package actors

type ActorData struct {
	ID        int64
	Name      string
	Gender    string
	BirthDate string
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
