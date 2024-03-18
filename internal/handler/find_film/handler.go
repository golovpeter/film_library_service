package find_film

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

// @Description	 Find film by title or actor
// @Tags         Films
// @Accept       json
// @Produce      json
// @Param search_field query string false "Search field"
// @Param value query string false "Searched value"
// @Param Authorization header string true "Bearer <token>" default("")
// @Success 200 {object} FilmData
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /film/find [get]
func (h *handler) FindFilm(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	searchField := queryValues.Get("search_field")
	searchValue := queryValues.Get("value")

	valid, errMsg := validateQueryParam(searchField)
	if !valid {
		h.logger.Error(errMsg)
		common.MakeErrorResponse(w, http.StatusBadRequest, errors.New(errMsg))
		return
	}

	serviceFilm, err := h.service.FindFilm(r.Context(), &films.FindFilmIn{
		SearchField: searchField,
		Value:       searchValue,
	})
	if err != nil {
		h.logger.WithError(err).Error(common.CreateActorError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.GettingFilmsError)
		return
	}

	jsonOut, err := json.Marshal(&FilmData{
		ID:          serviceFilm.ID,
		Title:       serviceFilm.Title,
		Description: serviceFilm.Description,
		ReleaseDate: serviceFilm.ReleaseDate,
		Rating:      serviceFilm.Rating,
	})
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
