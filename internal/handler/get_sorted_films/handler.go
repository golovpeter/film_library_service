package get_sorted_films

import (
	"encoding/json"
	"errors"
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

// GettingFilms godoc
// @Description	 Get all sorted films
// @Tags         Films
// @Produce      json
// @Param Authorization header string true "Bearer <token>" default("")
// @Param order_by query string false "Sorted field"
// @Success 200 {object} []FilmData
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /films [get]
func (h *handler) GettingFilms(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	orderBy := queryValues.Get("order_by")

	valid, errMsg := validateQueryParam(orderBy)
	if !valid {
		h.logger.Error(errMsg)
		common.MakeErrorResponse(w, http.StatusBadRequest, errors.New(errMsg))
		return
	}

	serviceFilms, err := h.service.GettingSortedFilms(r.Context(), orderBy)
	if err != nil {
		h.logger.WithError(err).Error(common.GettingFilmsError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.GettingFilmsError)
		return
	}

	out := make([]*FilmData, len(serviceFilms))
	for i, film := range serviceFilms {
		out[i] = &FilmData{
			ID:          film.ID,
			Title:       film.Title,
			Description: film.Description,
			Rating:      film.Rating,
			ReleaseDate: film.ReleaseDate,
			Actors:      film.Actors,
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
