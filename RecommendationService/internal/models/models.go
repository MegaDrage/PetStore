package models

import "time"

type PetProfile struct {
	PetID      int64   `json:"pet_id" bson:"pet_id"`
	Species    string   `json:"species" bson:"species"`
	Breed      string   `json:"breed" bson:"breed"`
	Age        int      `json:"age" bson:"age"`
	Allergies  []string `json:"allergies" bson:"allergies"`
	Conditions []string `json:"conditions" bson:"conditions"`
}

type LabResult struct {
	PetID     int64 `json:"pet_id" bson:"pet_id"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
	Metrics   map[string]float64 `json:"metrics" bson:"metrics"`
}

type Recommendation struct {
	PetID       int64 `json:"pet_id" bson:"pet_id"`
	Type        string    `json:"type" bson:"type"`
	Content     string    `json:"content" bson:"content"`
	Priority    int       `json:"priority" bson:"priority"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	ScheduledAt time.Time `json:"scheduled_at,omitempty" bson:"scheduled_at,omitempty"`
}