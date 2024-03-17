package actors

type ActorDataIn struct {
	Name      string
	Gender    string
	BirthDate string
}

type DeleteActorIn struct {
	ActorID int64
}
