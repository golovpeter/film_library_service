package main

import (
	"fmt"
	"net/http"

	"github.com/golovpeter/vk_intership_test_task/internal/common"
	"github.com/golovpeter/vk_intership_test_task/internal/config"
	"github.com/golovpeter/vk_intership_test_task/internal/handler/login_user"
	"github.com/golovpeter/vk_intership_test_task/internal/handler/register_user"
	"github.com/golovpeter/vk_intership_test_task/internal/repository/users"
	usersservice "github.com/golovpeter/vk_intership_test_task/internal/service/users"
	"github.com/sirupsen/logrus"
)

const (
	casbinModelPath  = "../../casbin/model.conf"
	casbinPolicyPath = "../../casbin/policy.csv"
)

func main() {
	logger := logrus.New()

	cfg, err := config.Parse()
	if err != nil {
		logger.Error("error to parse config file")
		return
	}

	level, err := logrus.ParseLevel(cfg.Logger.Level)
	if err != nil {
		logger.Error("error to parse logger level")
		return
	}

	logger.SetLevel(level)

	//enf, err := casbin.NewEnforcer(casbinModelPath, casbinPolicyPath)
	//if err != nil {
	//	logger.Error("error create enforcer: ", err.Error())
	//	return
	//}

	dbConn, err := common.CreateDbClient(cfg.Database)
	if err != nil {
		logger.Error("error create db connection: ", err.Error())
		return
	}

	mux := http.NewServeMux()

	//authMiddleware := authorization.NewMiddleware(logger, enf)

	usersRepository := users.NewRepository(dbConn)

	usersService := usersservice.NewService(usersRepository, cfg.Server.JWTKey)

	registerUserHandler := http.HandlerFunc(register_user.NewHandler(logger, usersService).Register)
	loginUserHandler := http.HandlerFunc(login_user.NewHandler(logger, usersService).Login)

	mux.Handle("POST /v1/user/register", registerUserHandler)
	mux.Handle("POST /v1/user/login", loginUserHandler)

	if err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), mux); err != nil {
		panic(err)
	}
}
