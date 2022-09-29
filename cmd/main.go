package main

import (
	"net/http"
	"strconv"

	"github.com/kalote/balance-server/pkg/conf"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	c      conf.Config
)

func main() {
	// Logger setup
	logger, _ = zap.NewProduction()
	defer logger.Sync()

	// Env vars: read & populate c
	err := envconfig.Process("", &c)
	if err != nil {
		logger.Fatal("Cannot read config", zap.Error(err))
	}

	// Router init
	r := NewRouter()

	// Server start
	port := strconv.Itoa(c.Port)
	logger.Info("Server running on port " + port)
	logger.Fatal(http.ListenAndServe(":"+port, r).Error())
}
