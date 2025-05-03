package main

import (
	"log"
	"net/http"

	"MegaDrage/PetStore/RecommendationService/internal/config"
	"MegaDrage/PetStore/RecommendationService/internal/repository"
	"MegaDrage/PetStore/RecommendationService/internal/api"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize repositories
	mongoRepo, err := repository.NewMongoRepository(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
	kafkaProducer, err := repository.NewKafkaProducer(cfg.KafkaBroker)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize service
	recService := service.NewRecommendationService(mongoRepo, kafkaProducer, cfg)

	// Initialize API
	router := api.NewRouter(recService)

	// Start server
	log.Printf("Starting Recommendation Service on :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
