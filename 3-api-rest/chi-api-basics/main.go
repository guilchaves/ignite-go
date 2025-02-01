package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/time", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Fprintln(w, now)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/users", func(w http.ResponseWriter, r *http.Request) {})
		})

		r.Route("/v2", func(r chi.Router) {})

        // use regex to only accept numbers as id param
		r.With(middleware.RealIP).
            Get("/users/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")
                fmt.Println(id)
			})

		r.Group(func(r chi.Router) {
			r.Use(middleware.BasicAuth("", map[string]string{
				"admin": "admin",
			}))

			// curl -u admin:admin http://127.0.0.1:8080/api/healthcheck
			r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "🦆 quack quack!")

			})
		})
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
