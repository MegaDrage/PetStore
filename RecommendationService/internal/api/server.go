package api

import (
	"github.com/gorilla/mux"
		
	"MegaDrage/PetStore/RecommendationService/internal/service"
)

func NewRouter(svc *service.RecommendationService) *mux.Router {
	router := mux.NewRouter()
	handler := NewHandler(svc)

	router.HandleFunc("/recommendations/{pet_id}", handler.GenerateRecommendations).Methods("GET")

	return router
}