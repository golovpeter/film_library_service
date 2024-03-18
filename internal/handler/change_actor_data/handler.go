package change_actor_data

import (
	"encoding/json"
	"errors"
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

// ChangeActorData godoc
// @Description	 Change actor data
// @Tags         Actors
// @Accept       json
// @Param request body ChangeActorDataIn true "request"
// @Param Authorization header string true "Bearer <token>" default("")
// @Success 200
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /actor/change [post]
func (h *handler) ChangeActorData(w http.ResponseWriter, r *http.Request) {
	var in *ChangeActorDataIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		h.logger.WithError(err).Error(common.BindJSONError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.BindJSONError)
		return
	}

	valid, errMsg := validateInParams(in)
	if !valid {
		h.logger.Error(errMsg)
		common.MakeErrorResponse(w, http.StatusBadRequest, errors.New(errMsg))
		return
	}

	err = h.service.ChangeActorInfo(r.Context(), &actors.ChangeActorDataIn{
		ID:        in.ID,
		Name:      in.Name,
		Gender:    in.Gender,
		BirthDate: in.BirthDate,
	})
	if err != nil {
		h.logger.WithError(err).Error(common.ChangeActorDataError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.ChangeActorDataError)
		return
	}
}
