package handler

import (
	"github.com/kalote/balance-server/pkg/conf"
	"go.uber.org/zap"
)

type Handler struct {
	c      *conf.Config
	logger *zap.Logger
}

type ProbeResponse struct {
	Status string `json:"status"`
	Type   string `json:"type"`
}

type BalanceResponse struct {
	Balance string `json:"balance"`
}

type InfuraResponse struct {
	Result string `json:"result"`
}

func NewHandler(c *conf.Config, logger *zap.Logger) *Handler {
	return &Handler{
		c,
		logger,
	}
}
