package core

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moritiza/gift-code/report-service/app/helper"
	"github.com/moritiza/gift-code/report-service/config"
)

// router create gorilla mux router and define routes
func Router(cfg config.Config) *mux.Router {
	r := mux.NewRouter()
	d := PrepareDependensies(cfg)

	// Create group routes
	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		helper.SuccessResponse(w, "ok", "pong", true, http.StatusOK)
	}).Methods(http.MethodGet)

	s.HandleFunc("/report", d.Handlers.ReportHandler.Get).Methods(http.MethodGet)
	s.HandleFunc("/report", d.Handlers.ReportHandler.Create).Methods(http.MethodPost)

	return r
}
