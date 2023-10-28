package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-go/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	_, err := database.GetPostgresDB("user=postgres password=password01 dbname=example sslmode=disable")
	if err != nil {
		log.Fatalf("error on connecting to database: %v", err)
	}

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]any{
			"status": "ok",
			"data":   "world",
		})
	})

	port := "3000"

	log.Printf("listening to port %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
