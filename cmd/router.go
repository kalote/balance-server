package main

import (
	"github.com/gorilla/mux"
	"github.com/kalote/balance-server/pkg/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	muxprom "gitlab.com/msvechla/mux-prometheus/pkg/middleware"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	instrumentation := muxprom.NewDefaultInstrumentation()
	router.Use(instrumentation.Middleware)

	router.Path("/metrics").Handler(promhttp.Handler())

	h := handler.NewHandler(logger)

	ethRouter := router.PathPrefix("/eth").Subrouter()
	ethRouter.HandleFunc("/balance/{ethAddress}", h.GetBalance)

	probesRouter := router.PathPrefix("/_meta").Subrouter()
	probesRouter.HandleFunc("/live", h.LiveProbe)
	probesRouter.HandleFunc("/ready", h.ReadyProbe)
	return router
}
