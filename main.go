package main

import (
	"log"
	"net/http"

	"github.com/daalfox/medieval_tours/internal/tour"
	"github.com/go-chi/chi/v5"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}

}

func run() error {
	r := chi.NewRouter()

	tourRepo := tour.NewPgRepo(nil)
	tourService := tour.NewTourService(&tourRepo)

	r.Post("/tours", tour.PostTourHandler(tourService))

	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return s.ListenAndServe()
}
