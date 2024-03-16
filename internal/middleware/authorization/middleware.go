package authorization

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/sirupsen/logrus"
)

type middleware struct {
	logger   *logrus.Logger
	enforcer *casbin.Enforcer
}

func NewMiddleware(
	logger *logrus.Logger,
	enforcer *casbin.Enforcer,
) *middleware {
	return &middleware{
		logger:   logger,
		enforcer: enforcer,
	}
}

func (m *middleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("not imp middleware")
	})
}
