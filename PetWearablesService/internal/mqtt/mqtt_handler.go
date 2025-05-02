package mqtt

import (
	"encoding/json"		

	"github.com/MegaDrage/PetStore/PetWearablesService/internal/models"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/storage"
	"github.com/MegaDrage/PetStore/PetWearablesService/pkg/logger"
)

type MqttHandler struct {
	influxClient *storage.InfluxClient
	logger       *logger.Logger
}

func NewMqttHandler(influxClient *storage.InfluxClient, logger *logger.Logger) *MqttHandler {
	return &MqttHandler{
		influxClient: influxClient,
		logger:       logger,
	}
}

func (h *MqttHandler) Handle(payload []byte) {
	var metrics models.CollarMetrics
	if err := json.Unmarshal(payload, &metrics); err != nil {
		h.logger.Error("Failed to parse MQTT payload", "error", err)
		return
	}

	data := models.CollarMetrics{
		PetID:       metrics.PetID,
		CollarID:    metrics.CollarID,
		Temperature: metrics.Temperature,
		HeartRate:   metrics.HeartRate,
		Location: struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		}{
			Lat: metrics.Location.Lat,
			Lon: metrics.Location.Lon,
		},
		Timestamp: metrics.Timestamp,
	}

	if err := h.influxClient.Save(data); err != nil {
		h.logger.Error("Failed to save MQTT data", "error", err, "pet_id", metrics.CollarID)
		return
	}
}