package models

import "time"

type Activity struct {
	Name        string    `json:"id"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}
