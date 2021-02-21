package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

func main() {
	db := dbInit()
	_ = db
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/urls", getURLS)
	r.Post("/newurl", newURL)
	http.ListenAndServe(":3000", r)
}
