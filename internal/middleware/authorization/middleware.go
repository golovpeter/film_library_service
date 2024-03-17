package authorization

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/sirupsen/logrus"
)

func AuthorizationMiddleware(logger *logrus.Logger, enforcer *casbin.Enforcer, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
