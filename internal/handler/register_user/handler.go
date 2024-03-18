package register_user

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

// Register godoc
// @Description	 Register in service
// @Tags         Users
// @Accept       json
// @Param request body UserDataIn true "request"
// @Success 200
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /user/register [post]
func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	var in UserDataIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		h.logger.WithError(err).Error(common.BindJSONError)
		common.MakeErrorResponse(w, http.StatusBadRequest, common.BindJSONError)
		return
	}

	passwordHash := common.GeneratePasswordHash(in.Password)

	err = h.service.Register(r.Context(), &users.UserDataIn{
		Username: in.Username,
		Password: passwordHash,
	})

	if err != nil {
		h.logger.WithError(err).Error(common.RegisterUserError)
		common.MakeErrorResponse(w, http.StatusInternalServerError, common.RegisterUserError)
		return
	}
}
