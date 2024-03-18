package login_user

import (
	"encoding/json"
	"net/http"

	"github.com/golovpeter/vk_intership_test_task/internal/common"
	"github.com/golovpeter/vk_intership_test_task/internal/service/users"
	"github.com/sirupsen/logrus"
)

type handler struct {
	logger *logrus.Logger

	service users.UserService
}

func NewHandler(
	logger *logrus.Logger,
	service users.UserService,
) *handler {
	return &handler{
		logger:  logger,
		service: service,
	}
}

// Login godoc
// @Description	 Login in service
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param request body UserDataIn true "request"
// @Success 200 {object} UserDataOut
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /user/login [post]
func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var in UserDataIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		h.logger.WithError(err).Error(common.BindJSONError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.BindJSONError)
		return
	}

	accessToken, err := h.service.Login(r.Context(), &users.UserDataIn{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		h.logger.WithError(err).Error(common.LoginUserError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.LoginUserError)
		return
	}

	out, err := json.Marshal(&UserDataOut{
		AccessToken: accessToken,
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
