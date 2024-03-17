package main

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/golovpeter/vk_intership_test_task/internal/common"
	"github.com/golovpeter/vk_intership_test_task/internal/config"
	"github.com/golovpeter/vk_intership_test_task/internal/handler/change_actor_data"
	"github.com/golovpeter/vk_intership_test_task/internal/handler/create_actor"
	"github.com/golovpeter/vk_intership_test_task/internal/handler/delete_actor"
	"github.com/golovpeter/vk_intership_test_task/internal/handler/login_user"
	"github.com/golovpeter/vk_intership_test_task/internal/handler/register_user"
	"github.com/golovpeter/vk_intership_test_task/internal/middleware/authorization"
	"github.com/golovpeter/vk_intership_test_task/internal/repository/actors"
	"github.com/golovpeter/vk_intership_test_task/internal/repository/users"
	actorsservice "github.com/golovpeter/vk_intership_test_task/internal/service/actors"
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

	enf, err := casbin.NewEnforcer(casbinModelPath, casbinPolicyPath)
	if err != nil {
		logger.Error("error create enforcer: ", err.Error())
		return
	}

	dbConn, err := common.CreateDbClient(cfg.Database)
	if err != nil {
		logger.Error("error create db connection: ", err.Error())
		return
	}

	r := http.NewServeMux()

	usersRepository := users.NewRepository(dbConn)
	actorsRepository := actors.NewRepository(dbConn)

	usersService := usersservice.NewService(usersRepository, cfg.Server.JWTKey)
	actorsService := actorsservice.NewService(actorsRepository)

	registerUserHandler := register_user.NewHandler(logger, usersService)
	loginUserHandler := login_user.NewHandler(logger, usersService)
	createActorHandler := create_actor.NewHandler(logger, actorsService)
	changeActorDataHandler := change_actor_data.NewHandler(logger, actorsService)
	deleteActorHandler := delete_actor.NewHandler(logger, actorsService)

	r.HandleFunc("POST /v1/user/register", registerUserHandler.Register)
	r.HandleFunc("POST /v1/user/login", loginUserHandler.Login)
	r.HandleFunc("POST /v1/actor/create", createActorHandler.CreateActor)
	r.HandleFunc("POST /v1/actor/change", changeActorDataHandler.ChangeActorData)
	r.HandleFunc("DELETE /v1/actor/delete", deleteActorHandler.DeleteActor)

	mux := authorization.AuthorizationMiddleware(logger, enf, r)

	if err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), mux); err != nil {
		panic(err)
	}
}
