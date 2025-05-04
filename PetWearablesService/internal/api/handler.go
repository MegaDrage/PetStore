package api

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"time"
	"fmt"

	"github.com/MegaDrage/PetStore/PetWearablesService/internal/models"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/mqtt"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/storage"
	"github.com/MegaDrage/PetStore/PetWearablesService/pkg/logger"
	"github.com/google/uuid"
)

type Handler struct {
	influxClient    *storage.InfluxClient
	mqttClient mqtt.Client
	logger     *logger.Logger
}

func NewHandler(influxClient *storage.InfluxClient, mqttClient mqtt.Client, logger *logger.Logger) *Handler {
	return &Handler{
		influxClient:   influxClient,
		mqttClient: mqttClient,
		logger:     logger,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/pets/wearables/{pet_id}/metrics", h.GetPetMetrics)
	mux.HandleFunc("POST /api/pets/wearables/simulate", h.SimulateMetrics)
}

func (h *Handler) GetPetMetrics(w http.ResponseWriter, r *http.Request) {
	petID := r.PathValue("pet_id")

	durationStr := r.URL.Query().Get("duration")

	var err error

	if _, err = uuid.Parse(petID); err != nil {
		h.logger.Error("Invalid pet_id: ", petID, ", error: ", err)
		http.Error(w, "Invalid pet_id: must be a valid UUID", http.StatusBadRequest)
		return
	}

	duration := 0 * time.Minute
	if durationStr != "" {
		duration, err = time.ParseDuration(durationStr)
		if err != nil {
			h.logger.Error("Invalid duration format, ", "error: ", err, ", duration: ", durationStr)
			http.Error(w, "Invalid duration format (examples: 15m, 1h, 30s)", http.StatusBadRequest)
			return
		}
	}

	h.logger.Debug("Processing request for pet metrics", "pet_id", petID)

	metrics, err := h.influxClient.GetMetrics(r.Context(), petID, duration)
	if err != nil {
		h.logger.Error("Failed to get metrics, ", "error: ", err, ", pet_id: ", petID)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.logger.Debug("Returning metrics response, ", "pet_id: ", petID, ", metrics_count: ", len(metrics))

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		h.logger.Error("Failed to encode response, ", "error: ", err, ", pet_id: ", petID)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SimulateMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		PetID string `json:"pet_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to parse request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	petID, err := uuid.Parse(req.PetID)
	if err != nil {
		h.logger.Error("Invalid UUID", "pet_id", req.PetID, "error", err)
		http.Error(w, "Invalid pet_id: must be a valid UUID", http.StatusBadRequest)
		return
	}


	metrics := models.CollarMetrics{
		PetID:       petID.String(),
		Temperature: 37.5 + rand.Float64()*3,
		HeartRate:   80 + rand.Intn(60),
		Location: models.Location {
			Lat: 55.7558 + (rand.Float64()-0.5)*0.5,
			Lon: 37.6173 + (rand.Float64()-0.5)*0.5,
		},
		Timestamp: time.Now().UTC(),
	}

	if metrics.PetID == "" || metrics.Temperature == 0 || metrics.HeartRate == 0 || metrics.Location.Lat == 0 || metrics.Location.Lon == 0 {
		h.logger.Error("Generated invalid metrics", "pet_id", petID.String(), "metrics", fmt.Sprintf("%+v", metrics))
		http.Error(w, "Failed to generate valid metrics", http.StatusInternalServerError)
		return
	}

	if err := h.influxClient.Save(metrics); err != nil {
		h.logger.Error("Failed to save metrics", "pet_id", petID.String(), "error", err)
		http.Error(w, "Failed to save metrics", http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(metrics)
	if err != nil {
		h.logger.Error("Error marshaling JSON", "pet_id", petID.String(), "error", err)
		http.Error(w, "Failed to marshal metrics", http.StatusInternalServerError)
		return
	}

	topic := fmt.Sprintf("pet/wearables/%s/metrics", petID.String())
	if err := h.mqttClient.Publish(topic, payload); err != nil {
		h.logger.Error("Error publishing to MQTT", "topic", topic, "error", err)
		http.Error(w, "Failed to publish metrics", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Published metrics to MQTT: ", topic, ", payload: ", string(payload))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"pet_id": petID.String(),
	})
}
