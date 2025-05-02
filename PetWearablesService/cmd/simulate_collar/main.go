package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type CollarMetrics struct {
	PetID       int64     `json:"pet_id"`
	CollarID    string    `json:"collar_id"`
	Temperature float64   `json:"temperature"`
	HeartRate   int       `json:"heart_rate"`
	Location    struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
	Timestamp time.Time `json:"timestamp"`
}

type Config struct {
	Broker   string
	ClientID string
	Username string
	Password string
}

func LoadConfig() (Config, error) {
	cfg := Config{
		Broker:   os.Getenv("MQTT_BROKER"),
		ClientID: os.Getenv("MQTT_CLIENT_ID"),
		Username: os.Getenv("MQTT_USERNAME"),
		Password: os.Getenv("MQTT_PASSWORD"),
	}

	if cfg.Broker == "" {
		return cfg, fmt.Errorf("MQTT_BROKER is not set")
	}
	if cfg.ClientID == "" {
		return cfg, fmt.Errorf("MQTT_CLIENT_ID is not set")
	}

	return cfg, nil
}

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	opts := mqtt.NewClientOptions().
		AddBroker(cfg.Broker).
		SetClientID(cfg.ClientID + "-simulator")
	if cfg.Username != "" && cfg.Password != "" {
		opts.SetUsername(cfg.Username)
		opts.SetPassword(cfg.Password)
	}

	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(10 * time.Second)
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		fmt.Println("Connected to MQTT broker")
	})
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v\n", err)
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Error connecting to MQTT: %v\n", token.Error())
		os.Exit(1)
	}
	defer client.Disconnect(250)

	pets := []struct {
		PetID    int64
		CollarID string
	}{
		{PetID: 123, CollarID: "collar1"},
		{PetID: 456, CollarID: "collar2"},
		{PetID: 789, CollarID: "collar3"},
	}

	for {
		for _, pet := range pets {
			metrics := CollarMetrics{
				PetID:       pet.PetID,
				CollarID:    pet.CollarID,
				Temperature: 37.0 + rand.Float64()*2.0,
				HeartRate:   90 + rand.Intn(40),
				Location: struct {
					Lat float64 `json:"lat"`
					Lon float64 `json:"lon"`
				}{
					Lat: 55.7558 + (rand.Float64()-0.25)*0.04,
					Lon: 37.6173 + (rand.Float64()-0.5)*0.02,
				},
				Timestamp: time.Now().UTC(),
			}

			payload, err := json.Marshal(metrics)
			if err != nil {
				fmt.Printf("Error marshaling JSON for %s: %v\n", pet.CollarID, err)
				continue
			}

			topic := fmt.Sprintf("pet/wearables/%s/metrics", pet.CollarID)
			if token := client.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
				fmt.Printf("Error publishing to %s: %v\n", topic, token.Error())
			} else {
				fmt.Printf("Published to %s: %s\n", topic, string(payload))
			}
		}

		time.Sleep(60 * time.Second)
	}
}
