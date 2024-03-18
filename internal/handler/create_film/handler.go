package create_film

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

func (h *handler) CreateFilm(w http.ResponseWriter, r *http.Request) {
	var in *CreatFilmIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		h.logger.WithError(err).Error(common.BindJSONError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.BindJSONError)
		return
	}

	err = h.service.CreateFilm(r.Context(), &films.FilmData{
		Title:       in.Title,
		Description: in.Description,
		ReleaseDate: in.ReleaseDate,
		Rating:      in.Rating,
		Actors:      in.Actors,
	})
	if err != nil {
		h.logger.WithError(err).Error(common.CreateActorError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.CreateFilmError)
		return
	}
}
