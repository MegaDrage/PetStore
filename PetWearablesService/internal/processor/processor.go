package processor

import (
	"encoding/json"

	"github.com/MegaDrage/PetStore/PetWearablesService/internal/config"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/models"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/storage"
	"github.com/MegaDrage/PetStore/PetWearablesService/pkg/logger"
)

type Processor struct {
	store  *storage.InfluxClient
	cfg    *config.Config
	logger *logger.Logger
}

func NewProcessor(store *storage.InfluxClient, cfg *config.Config, logger *logger.Logger) *Processor {
	return &Processor{store: store, cfg: cfg, logger: logger}
}

func (p *Processor) Process(payload []byte) {
	var data models.PetData
	if err := json.Unmarshal(payload, &data); err != nil {
		p.logger.WithFields(map[string]interface{}{
			"error": err.Error(),
		}).Error("Error decoding data")
		return
	}

	p.logger.WithFields(map[string]interface{}{
		"pet_id": data.PetID,
	}).Info("Getting from devices")

	if err := p.store.Save(data); err != nil {
		p.logger.WithFields(map[string]interface{}{
			"pet_id": data.PetID,
			"error":  err.Error(),
		}).Error("Error saving data")
	}
}