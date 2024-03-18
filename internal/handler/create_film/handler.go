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

// CreateFilm godoc
// @Description	 Insert new film
// @Tags         Films
// @Accept       json
// @Produce      json
// @Param request body CreatFilmIn true "request"
// @Param Authorization header string true "Bearer <token>" default("")
// @Success 200 {object} CreateFilmOut
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /film/create [post]
func (h *handler) CreateFilm(w http.ResponseWriter, r *http.Request) {
	var in *CreatFilmIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		h.logger.WithError(err).Error(common.BindJSONError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.BindJSONError)
		return
	}

	id, err := h.service.CreateFilm(r.Context(), &films.FilmData{
		Title:       in.Title,
		Description: in.Description,
		ReleaseDate: in.ReleaseDate,
		Rating:      in.Rating,
		Actors:      in.Actors,
	})
	if err != nil {
		h.logger.WithError(err).Error(common.CreateFilmError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.CreateFilmError)
		return
	}

	out, err := json.Marshal(&CreateFilmOut{
		ID: id,
	})
	if err != nil {
		h.logger.WithError(err).Error(err.Error())
		common.MakeErrorResponse(w, http.StatusBadRequest, common.BindJSONError)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		h.logger.WithError(err).Error(err.Error())
		common.MakeErrorResponse(w, http.StatusBadRequest, err)
		return
	}
}
