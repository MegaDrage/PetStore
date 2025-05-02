package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"


	"github.com/MegaDrage/PetStore/PetWearablesService/internal/storage"
	"github.com/MegaDrage/PetStore/PetWearablesService/pkg/logger"
	"github.com/gorilla/mux"
)

type Handler struct {
	influxClient *storage.InfluxClient
	logger       *logger.Logger
}

func NewHandler(influxClient *storage.InfluxClient, logger *logger.Logger) *Handler {
	return &Handler{influxClient: influxClient, logger: logger}
}

func (h *Handler) GetPetMetrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petIDStr := vars["pet_id"]
	durationStr := r.URL.Query().Get("duration")

	petID, err := strconv.ParseInt(petIDStr, 10, 64)
	if err != nil {
		h.logger.Error("Invalid pet_id format, ", "error: ", err, ", pet_id:", petIDStr)
		http.Error(w, "Invalid pet_id format", http.StatusBadRequest)
		return
	}

	duration := 15 * time.Minute
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
