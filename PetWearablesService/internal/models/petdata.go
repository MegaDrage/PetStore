package models

type PetData struct {
	PetID       string  `json:"pet_id"`
	Temperature float64 `json:"temperature"`
	HeartRate   int     `json:"heart_rate"`
	Location    string  `json:"location"`
}