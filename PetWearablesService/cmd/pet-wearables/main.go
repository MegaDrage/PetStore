package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/MegaDrage/PetStore/PetWearablesService/internal/config"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/mqtt"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/processor"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/storage"
	"github.com/MegaDrage/PetStore/PetWearablesService/pkg/logger"
)

func main() {
	logger := logger.NewLogger()

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Error loading configuration:", err)
	}

	store, err := storage.NewInfluxClient(cfg.InfluxDB)
	if err != nil {
		logger.Fatal("Error connecting InfluxDB:", err)
	}
	defer store.Close()

	proc := processor.NewProcessor(store, cfg, logger)

	mqttClient, err := mqtt.NewClient(cfg.MQTT, proc.Process, logger)
	if err != nil {
		logger.Fatal("Error intializing MQTT-client:", err)
	}
	defer mqttClient.Disconnect(250)

	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		logger.Info("Signal for graceful shutdown")
		cancel()
	}()

	<-ctx.Done()
	logger.Info("Service Stopped Working")
}