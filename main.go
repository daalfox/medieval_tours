package main

import (
	"log"
	"net/http"

	"github.com/daalfox/medieval_tours/internal/schedule"
	"github.com/daalfox/medieval_tours/internal/tour"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}

}

func run() error {
	r := chi.NewRouter()

	var pool pgxpool.Pool

	tourRepo := tour.NewPgRepo(&pool)
	tourService := tour.NewTourService(&tourRepo)
	scheduleRepo := schedule.NewPgRepo(&pool)
	scheduleService := schedule.NewScheduleService(&scheduleRepo)

	r.Post("/tours", tour.PostTourHandler(tourService))
	r.Post("/schedules", schedule.PostScheduleHandler(scheduleService))

	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return s.ListenAndServe()
}
