package main

import (
	"encoding/json"
	"fmt"
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

	r.Post("/tours", func(w http.ResponseWriter, r *http.Request) {
		var payload tour.Tour

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(nil)
			return
		}

		fmt.Println(payload)

		tourService.Insert(r.Context(), payload)
	})

	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return s.ListenAndServe()
}
