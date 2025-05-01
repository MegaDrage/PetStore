package storage

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/config"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/models"
)

type InfluxClient struct {
	client   influxdb2.Client
	org      string
	bucket   string
}

func NewInfluxClient(cfg config.InfluxDBConfig) (*InfluxClient, error) {
	client := influxdb2.NewClient(cfg.URL, cfg.Token)
	return &InfluxClient{
		client: client,
		org:    cfg.Org,
		bucket: cfg.Bucket,
	}, nil
}

func (c *InfluxClient) Save(data models.PetData) error {
	writeAPI := c.client.WriteAPIBlocking(c.org, c.bucket)

	point := influxdb2.NewPointWithMeasurement("pet_data").
		AddTag("pet_id", data.PetID).
		AddField("temperature", data.Temperature).
		AddField("heart_rate", data.HeartRate).
		SetTime(time.Now())

	return writeAPI.WritePoint(context.Background(), point)
}

func (c *InfluxClient) GetMetrics(ctx context.Context, petID string) ([]models.PetData, error) {
	queryAPI := c.client.QueryAPI(c.org)
	query := `
		from(bucket: "` + c.bucket + `")
			|> range(start: -1h)
			|> filter(fn: (r) => r._measurement == "pet_data")
			|> filter(fn: (r) => r.pet_id == "` + petID + `")
	`

	result, err := queryAPI.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var metrics []models.PetData
	for result.Next() {
		record := result.Record()
		metric := models.PetData{
			PetID:       petID,
			Temperature: record.ValueByKey("temperature").(float64),
			HeartRate:   int(record.ValueByKey("heart_rate").(float64)),
			Location:    "unknown",
		}
		metrics = append(metrics, metric)
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	return metrics, nil
}

func (c *InfluxClient) Close() {
	c.client.Close()
}