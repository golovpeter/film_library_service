package actors

import "context"

type ActorService interface {
	CreateActor(ctx context.Context, data *ActorDataIn) error
	ChangeActorInfo(ctx context.Context, data *ActorDataIn) error
	DeleteActor(ctx context.Context, data *DeleteActorIn) error
}
