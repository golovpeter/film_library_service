package main

import (
	"fmt"
	"net/http"

	"github.com/golovpeter/vk_intership_test_task/internal/config"
	"github.com/sirupsen/logrus"
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

	mux := http.NewServeMux()

	if err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), mux); err != nil {
		panic(err)
	}
}
