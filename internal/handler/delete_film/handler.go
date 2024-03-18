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

// DeleteFilm godoc
// @Description	 Delete all data about film
// @Tags         Films
// @Accept       json
// @Param request body DeleteFilmIn true "request"
// @Param Authorization header string true "Bearer <token>" default("")
// @Success 200
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /film/delete [delete]
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
		h.logger.WithError(err).Error(common.DeleteFilmError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.DeleteFilmError)
		return
	}
}
