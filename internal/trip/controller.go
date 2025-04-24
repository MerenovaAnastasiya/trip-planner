package trip

import (
	"encoding/json"
	"github.com/MerenovaAnastasiya/trip-planner/pkg/models"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func GetAllTrips(w http.ResponseWriter, r *http.Request) {
	trips := []models.Trip{
		{
			ID:          uuid.New(),
			Destination: models.Country{Name: "Испания"},
			StartDate:   time.Date(2025, 7, 10, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 7, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:          uuid.New(),
			Destination: models.Country{Name: "Исландия"},
			StartDate:   time.Date(2025, 9, 5, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 9, 15, 0, 0, 0, 0, time.UTC),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trips)
}
