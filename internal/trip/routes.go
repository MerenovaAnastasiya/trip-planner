package trip

import "github.com/go-chi/chi/v5"

func RegisterRoutes(r chi.Router) {
	r.Route("/trips", func(r chi.Router) {
		r.Get("/", GetAllTrips)
	})
}
