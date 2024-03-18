package actors

import (
	"context"

	"github.com/golovpeter/vk_intership_test_task/internal/repository/actors"
)

type service struct {
	repository actors.Repository
}

func NewService(repository actors.Repository) ActorService {
	return &service{repository: repository}
}

func (s *service) CreateActor(ctx context.Context, data *ActorData) (int64, error) {
	return s.repository.CreateActor(ctx, &actors.ActorData{
		Name:      data.Name,
		Gender:    data.Gender,
		BirthDate: data.BirthDate,
	})
}

func (s *service) ChangeActorInfo(ctx context.Context, data *ChangeActorDataIn) error {
	return s.repository.ChangeActorInfo(ctx, &actors.ChangeActorDataIn{
		ID:        data.ID,
		Name:      data.Name,
		Gender:    data.Gender,
		BirthDate: data.BirthDate,
	})
}

func (s *service) DeleteActor(ctx context.Context, data *DeleteActorIn) error {
	return s.repository.DeleteActor(ctx, &actors.DeleteActorIn{
		ActorID: data.ActorID,
	})
}

func (s *service) GetAllActors(ctx context.Context) ([]*ActorData, error) {
	repoActors, err := s.repository.GetAllActors(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]*ActorData, len(repoActors))
	for i, actor := range repoActors {
		out[i] = &ActorData{
			ID:        actor.ID,
			Name:      actor.Name,
			Gender:    actor.Gender,
			BirthDate: actor.BirthDate,
			Films:     actor.Films,
		}
	}

	return out, nil
}
