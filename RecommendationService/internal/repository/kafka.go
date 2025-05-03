package repository

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"MegaDrage/PetStore/RecommendationService/internal/models"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(broker string) (*KafkaProducer, error) {
	producer, err := sarama.NewSyncProducer([]string{broker}, nil)
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{producer: producer}, nil
}

func (p *KafkaProducer) PublishNotification(rec models.Recommendation) error {
	msg, _ := json.Marshal(rec)
	_, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "notifications",
		Value: sarama.StringEncoder(msg),
	})
	return err
}