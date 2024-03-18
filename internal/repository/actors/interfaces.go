package actors

import "context"

type Repository interface {
	CreateActor(ctx context.Context, data *ActorData) error
	ChangeActorInfo(ctx context.Context, data *ChangeActorDataIn) error
	DeleteActor(ctx context.Context, data *DeleteActorIn) error
	GetAllActors(ctx context.Context) ([]*ActorData, error)
}
