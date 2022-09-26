package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/muhaidil13/booking/internal/config"
	"github.com/muhaidil13/booking/internal/handlers"
)

func routes(App *config.Appconfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/room", handlers.Repo.Room)

	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Post("/reservation", handlers.Repo.PostReservation)
	mux.Post("/reservation-json", handlers.Repo.AvailabilityJson)

	mux.Get("/reservation-result", handlers.Repo.ReservationSummary)
	mux.Post("/make-reservation", handlers.Repo.PostMake_Reservation)
	mux.Get("/make-reservation", handlers.Repo.Make_Reservation)
	mux.Get("/pemesannan", handlers.Repo.Pemesannan)

	fileserver := http.FileServer(http.Dir("../static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))
	return mux
}
