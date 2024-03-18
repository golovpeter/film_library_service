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

// ChangeFilmData godoc
// @Description	 Change film data
// @Tags         Films
// @Accept       json
// @Param request body ChangeFilmIn true "request"
// @Param Authorization header string true "Bearer <token>" default("")
// @Success 200
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /film/change [post]
func (h *handler) ChangeFilmData(w http.ResponseWriter, r *http.Request) {
	var in *ChangeFilmIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		h.logger.WithError(err).Error(common.BindJSONError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.BindJSONError)
		return
	}

	err = h.service.ChangeFilmData(r.Context(), &films.FilmData{
		ID:          in.ID,
		Title:       in.Title,
		Description: in.Description,
		Rating:      in.Rating,
		ReleaseDate: in.ReleaseDate,
		Actors:      in.Actors,
	})
	if err != nil {
		h.logger.WithError(err).Error(common.ChangeActorDataError)
		common.MakeErrorResponse(w, http.StatusInternalServerError, common.ChangeFilmDataError)
		return
	}
}
