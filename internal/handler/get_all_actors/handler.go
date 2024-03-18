package get_all_actors

import (
	"encoding/json"
	"net/http"

	"github.com/golovpeter/vk_intership_test_task/internal/common"
	"github.com/golovpeter/vk_intership_test_task/internal/service/actors"
	"github.com/sirupsen/logrus"
)

type handler struct {
	logger *logrus.Logger

	service actors.ActorService
}

func NewHandler(
	logger *logrus.Logger,
	service actors.ActorService,
) *handler {
	return &handler{
		logger:  logger,
		service: service,
	}
}

func (h *handler) GetAllActors(w http.ResponseWriter, r *http.Request) {
	serviceActors, err := h.service.GetAllActors(r.Context())
	if err != nil {
		h.logger.WithError(err).Error(common.CreateActorError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.GettingFilmsError)
		return
	}

	out := make([]*ActorData, len(serviceActors))
	for i, actor := range serviceActors {
		out[i] = &ActorData{
			ID:        actor.ID,
			Name:      actor.Name,
			Gender:    actor.Gender,
			BirthDate: actor.BirthDate,
			Films:     actor.Films,
		}
	}

	jsonOut, err := json.Marshal(out)
	if err != nil {
		h.logger.WithError(err).Error(err.Error())
		common.MakeErrorResponse(w, http.StatusBadRequest, common.BindJSONError)
		return
	}

	_, err = w.Write(jsonOut)
	if err != nil {
		h.logger.WithError(err).Error(err.Error())
		common.MakeErrorResponse(w, http.StatusBadRequest, err)
		return
	}
}
