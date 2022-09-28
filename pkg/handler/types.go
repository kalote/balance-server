package handler

import "go.uber.org/zap"

type handler struct {
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger) *handler {
	return &handler{
		logger,
	}
}
