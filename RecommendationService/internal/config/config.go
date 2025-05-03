package config

type Config struct {
	MongoURI      string
	KafkaBroker   string
	ProfileURL    string
	LabURL        string
	Port          string
}

func Load() Config {
	return Config{
		MongoURI:    "mongodb://localhost:27017",
		KafkaBroker: "localhost:9092",
		ProfileURL:  "http://profile-service:8080",
		LabURL:      "http://lab-service:8080",
		Port:        "8083",
	}
}