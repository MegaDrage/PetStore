package api

import (
	"encoding/json"
	"net/http"

	"MegaDrage/PetStore/RecommendationService/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.RecommendationService
}

func NewHandler(svc *service.RecommendationService) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) GenerateRecommendations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petID := vars["pet_id"]

	recommendations, err := h.service.GenerateRecommendations(r.Context(), petID)
	if err != nil {
		http.Error(w, "Failed to generate recommendations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}
