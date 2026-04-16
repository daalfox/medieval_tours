package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}

}

func run() error {
	r := chi.NewRouter()

	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return s.ListenAndServe()
}
