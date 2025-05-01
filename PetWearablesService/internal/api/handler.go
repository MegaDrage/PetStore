package api

import (
	"encoding/json"
	"net/http"

	"github.com/MegaDrage/PetStore/PetWearablesService/internal/storage"
	"github.com/MegaDrage/PetStore/PetWearablesService/pkg/logger"
	"github.com/gorilla/mux"
)

type Handler struct {
	store  *storage.InfluxClient
	logger *logger.Logger
}

func NewHandler(store *storage.InfluxClient, logger *logger.Logger) *Handler {
	return &Handler{store: store, logger: logger}
}

func (h *Handler) GetPetMetrics(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petID := vars["pet_id"]

	metrics, err := h.store.GetMetrics(r.Context(), petID)
	if err != nil {
		h.logger.Error("Errog gettings metrics:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		h.logger.Error("Error codding answer:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
}