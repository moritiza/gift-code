package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/moritiza/gift-code/wallet-service/config"
	"github.com/moritiza/gift-code/wallet-service/core"
)

func main() {
	cfg := core.Bootstrap()
	defer config.DisconnectDatabase(cfg)

	srv := &http.Server{
		Handler: handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "HEAD", "POST"}),
			handlers.AllowedHeaders([]string{"Accept", "Origin", "Content-Type", "api_key", "Authorization"}),
			handlers.AllowedOrigins([]string{"*"}))((core.Router(cfg))),
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		cfg.Logger.Fatalf("Server error: ", err.Error())
	}
}
