package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error during json marshal", "error", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("error sending response", "error", err)
	}
}

type User struct {
	Username string
	ID       int64 `json:"id,string"`
	Role     string
	Password Password `json:"-"`
}

type Password string

func (p Password) String() string {
	return "[REDACTED]"
}

func (p Password) LogValue() slog.Value {
	return slog.StringValue("[REDACTED]")
}

const LevelFoo = slog.Level(-50)

func main() {
	p := Password("123456")
	u := User{Password: p}
	slog.Info("password", "p", u)
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     LevelFoo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "level" {
				level := a.Value.String()
				if level == "DEBUG-46" {
					a.Value = slog.StringValue("FOO")
				}
			}
			return a
		},
	}
	l := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(l)
	slog.Debug("foo")
	slog.Info("starting server", "version", "1.0.0", "port", "8080")
	l = l.With(slog.Group("app_info", slog.String("version", "1.0.0")))
	l.Info("this is a test", "user", u)
	l.LogAttrs(context.Background(), LevelFoo, "any message")
	l.LogAttrs(
		context.Background(), slog.LevelInfo, "http request received",
		slog.Group("http_data",
			slog.String("method", http.MethodGet),
			slog.Int("status", http.StatusOK),
		),
		slog.Duration("time_taken", time.Second),
	)
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	db := map[int64]User{
		1: {
			Username: "admin",
			Password: "admin",
			Role:     "admin",
			ID:       1,
		},
	}

	r.Group(func(r chi.Router) {
		r.Use(jsonMiddleware)
		r.Get("/users/{id:[0-9]+}", handleGetUsers(db))
		r.Post("/users", handlePostUsers(db))
	})

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func handleGetUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, _ := strconv.ParseInt(idStr, 10, 64)

		// this is a race condition
		user, ok := db[id]
		if !ok {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)

	}
}

func handlePostUsers(db map[int64]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 10_000)
		data, err := io.ReadAll(r.Body)
		if err != nil {
			var maxErr *http.MaxBytesError
			if errors.As(err, &maxErr) {
				sendJSON(w, Response{Error: "body too large"}, http.StatusRequestEntityTooLarge)
				return
			}
			slog.Error("error reading body", "error", err)
			sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		var user User
		if err := json.Unmarshal(data, &user); err != nil {
			sendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		// this is a race condition
		db[user.ID] = user

		w.WriteHeader(http.StatusCreated)
	}

}
