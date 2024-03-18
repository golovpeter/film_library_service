package delete_film

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

func (h *handler) DeleteFilm(w http.ResponseWriter, r *http.Request) {
	var in *DeleteFilmIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		h.logger.WithError(err).Error(common.BindJSONError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.BindJSONError)
		return
	}

	err = h.service.DeleteFilm(r.Context(), &films.DeleteFilmIn{
		FilmID: in.FilmID,
	})
	if err != nil {
		h.logger.WithError(err).Error(common.CreateActorError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.DeleteFilmError)
		return
	}
}
