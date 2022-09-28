package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	muxprom "gitlab.com/msvechla/mux-prometheus/pkg/middleware"
)

func newRouter() *mux.Router {
	router := mux.NewRouter()
	instrumentation := muxprom.NewDefaultInstrumentation()
	router.Use(instrumentation.Middleware)

	router.Path("/metrics").Handler(promhttp.Handler())

	sub := router.PathPrefix("/eth").Subrouter()
	sub.HandleFunc("/balance/{ethAddress}", GetBalance)
	return router
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(vars["ethAddress"]))
}
