package mqtt

import (
	"time"

	"github.com/MegaDrage/PetStore/PetWearablesService/internal/config"
	"github.com/MegaDrage/PetStore/PetWearablesService/pkg/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	client mqtt.Client
	logger *logger.Logger
}

func NewClient(cfg config.MQTTConfig, handler func([]byte), logger *logger.Logger) (*Client, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(cfg.Broker).
		SetClientID(cfg.ClientID).
		SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
			handler(msg.Payload())
		})

	if cfg.Username != "" && cfg.Password != "" {
		opts.SetUsername(cfg.Username)
		opts.SetPassword(cfg.Password)
	}

	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(10 * time.Second)

	opts.SetOnConnectHandler(func(client mqtt.Client) {
		logger.Info("Success connect to MQTT-brocker")
		if token := client.Subscribe("pet/wearables/#", 0, nil); token.Wait() && token.Error() != nil {
			logger.Error("Resubscribe error:", token.Error())
		}
	})
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		logger.Warn("Connection with MQTT-brocker ended:", err)
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Error("Error connecting MQTT:", token.Error())
		return nil, token.Error()
	}

	if token := client.Subscribe("pet/wearables/#", 0, nil); token.Wait() && token.Error() != nil {
		logger.Error("Error subscribing:", token.Error())
		return nil, token.Error()
	}

	return &Client{client: client, logger: logger}, nil
}

func (c *Client) Disconnect(quiesce uint) {
	c.logger.Info("Disconnect from MQTT")
	c.client.Disconnect(quiesce)
}