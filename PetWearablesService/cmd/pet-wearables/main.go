package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MegaDrage/PetStore/PetWearablesService/internal/api"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/config"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/mqtt"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/storage"
	"github.com/MegaDrage/PetStore/PetWearablesService/pkg/logger"
)

func main() {
	logger := logger.NewLogger()

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load config", "error", err)
	}

	influxClient, err := storage.NewInfluxClient(cfg.InfluxDB, logger)
	if err != nil {
		logger.Fatal("Failed to initialize InfluxDB client", "error", err)
	}
	defer influxClient.Close()
	
	handler := mqtt.NewMqttHandler(influxClient, logger)
	mqttHandler := handler.Handle

	mqttClient, err := mqtt.NewClient(cfg.MQTT, mqttHandler, logger)
	if err != nil {
		logger.Fatal("Failed to initialize MQTT client, ", "error: ", err)
	}
	defer mqttClient.Disconnect(250)

	apiHandler := api.NewHandler(influxClient, *mqttClient, logger)
	server := api.NewServer(":8085", apiHandler, logger)

	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", "error", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Failed to shutdown server", "error", err)
	}
}
