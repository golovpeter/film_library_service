package change_film_data

import (
	"encoding/json"
	"net/http"

	"github.com/golovpeter/vk_intership_test_task/internal/common"
	"github.com/golovpeter/vk_intership_test_task/internal/service/films"
	"github.com/sirupsen/logrus"
)

type handler struct {
	logger *logrus.Logger

	service films.FilmsService
}

func NewHandler(
	logger *logrus.Logger,
	service films.FilmsService,
) *handler {
	return &handler{
		logger:  logger,
		service: service,
	}
}

func (h *handler) ChangeFilmData(w http.ResponseWriter, r *http.Request) {
	var in *ChangeFilmIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		h.logger.WithError(err).Error(common.BindJSONError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.BindJSONError)
		return
	}

	err = h.service.ChangeFilmData(r.Context(), &films.ChangeFilmIn{
		ID:          in.ID,
		Title:       in.Title,
		Description: in.Description,
		Rating:      in.Rating,
		ReleaseDate: in.ReleaseDate,
		Actors:      in.Actors,
	})
	if err != nil {
		h.logger.WithError(err).Error(common.ChangeActorDataError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.ChangeFilmDataError)
		return
	}
}
