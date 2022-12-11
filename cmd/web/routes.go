package main

import (
	"github.com/sagarshinganwade/booking_and_reservation/package/config"
	"github.com/sagarshinganwade/booking_and_reservation/package/handlers"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(r *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(LoadSession)
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/login", http.HandlerFunc(handlers.Repo.Login))

	return mux
}
