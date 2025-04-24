package main

import (
	"github.com/MerenovaAnastasiya/trip-planner/internal/trip"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	trip.RegisterRoutes(r)
	http.ListenAndServe(":8080", r)
}
