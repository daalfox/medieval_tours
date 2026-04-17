package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/daalfox/medieval_tours/internal/schedule"
	"github.com/daalfox/medieval_tours/internal/tour"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}

}

func run() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	r := chi.NewRouter()

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	tourRepo := tour.NewPgRepo(pool)
	tourService := tour.NewTourService(&tourRepo)
	scheduleRepo := schedule.NewPgRepo(pool)
	scheduleService := schedule.NewScheduleService(&scheduleRepo)

	r.Route("/tours", func(r chi.Router) {
		r.Get("/", tour.GetTourHandler(tourService))
		r.Post("/", tour.PostTourHandler(tourService))
	})
	r.Route("/schedules", func(r chi.Router) {
		r.Get("/", schedule.GetScheduleHandler(scheduleService))
		r.Post("/", schedule.PostScheduleHandler(scheduleService))
	})

	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return s.ListenAndServe()
}
