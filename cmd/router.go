package main

import (
	"github.com/gorilla/mux"
	"github.com/kalote/balance-server/pkg/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	muxprom "gitlab.com/msvechla/mux-prometheus/pkg/middleware"
)

func NewRouter() *mux.Router {
	// Init mux
	router := mux.NewRouter()

	// Init handler with config & logger
	h := handler.NewHandler(&c, logger)

	// Init prometheus metrics & mw & path
	instrumentation := muxprom.NewDefaultInstrumentation()
	router.Use(instrumentation.Middleware)
	router.Path("/metrics").Handler(promhttp.Handler())

	// Main route
	ethRouter := router.PathPrefix("/eth").Subrouter()
	ethRouter.HandleFunc("/balance/{ethAddress}", h.GetBalance)

	// Probes route
	probesRouter := router.PathPrefix("/_meta").Subrouter()
	probesRouter.HandleFunc("/healthz", h.LiveProbe)
	probesRouter.HandleFunc("/ready", h.ReadyProbe)

	return router
}
