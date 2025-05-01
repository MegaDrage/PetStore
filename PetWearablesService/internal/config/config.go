package config

import (
	"github.com/caarlos0/env/v11"
)

type InfluxDBConfig struct {
	URL    string `env:"INFLUXDB_URL" envDefault:"http://influxdb:8086"`
	Token  string `env:"INFLUXDB_TOKEN"`
	Org    string `env:"INFLUXDB_ORG" envDefault:"org"`
	Bucket string `env:"INFLUXDB_BUCKET" envDefault:"bucket"`
}

type MQTTConfig struct {
	Broker   string `env:"MQTT_BROKER" envDefault:"tcp://mosquitto:1883"`
	ClientID string `env:"MQTT_CLIENT_ID" envDefault:"pet-wearables-service"`
	Username string `env:"MQTT_USERNAME"`
	Password string `env:"MQTT_PASSWORD"`
}

type Config struct {
	MQTT              MQTTConfig
	InfluxDB          InfluxDBConfig
	ProfileServiceURL string `env:"PROFILE_SERVICE_URL" envDefault:"http://profile-service"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}