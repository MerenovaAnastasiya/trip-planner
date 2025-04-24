package models

import (
	"github.com/google/uuid"
	"time"
)

type Trip struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Destination Country   `json:"destination"`
}
