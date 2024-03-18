package delete_actor

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

// DeleteActor godoc
// @Description	 Delete actor
// @Tags         Actors
// @Accept       json
// @Param request body DeleteActorIn true "request"
// @Param Authorization header string true "Bearer <token>" default("")
// @Success 200
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /actor/delete [delete]
func (h *handler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	var in *DeleteActorIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		h.logger.WithError(err).Error(common.BindJSONError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.BindJSONError)
		return
	}

	err = h.service.DeleteActor(r.Context(), &actors.DeleteActorIn{
		ActorID: in.ActorID,
	})
	if err != nil {
		h.logger.WithError(err).Error(common.DeleteActorError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.DeleteActorError)
		return
	}
}
