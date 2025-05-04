package models

import (
	"time"

)

type CollarMetrics struct {
	PetID       string `json:"pet_id"`
	Temperature float64   `json:"temperature"`
	HeartRate   int       `json:"heart_rate"`
	Location    struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
	Timestamp time.Time `json:"timestamp"`
}