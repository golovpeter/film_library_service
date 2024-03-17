package actors

import "context"

type Repository interface {
	CreateActor(ctx context.Context, data *ActorDataIn) error
	ChangeActorInfo(ctx context.Context, data *ActorDataIn) error
	DeleteActor(ctx context.Context, data *DeleteActorIn) error
}
