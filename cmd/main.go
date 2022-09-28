package main

import (
	"net/http"

	"go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := newRouter()

	logger.Info("Server running on port 3000")
	logger.Fatal(http.ListenAndServe(":3000", r).Error())
}
