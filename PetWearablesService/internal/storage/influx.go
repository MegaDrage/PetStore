package storage

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/MegaDrage/PetStore/PetWearablesService/internal/config"
	"github.com/MegaDrage/PetStore/PetWearablesService/internal/models"
	"github.com/MegaDrage/PetStore/PetWearablesService/pkg/logger"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxClient struct {
	client influxdb2.Client
	org    string
	bucket string
	logger *logger.Logger
}

func NewInfluxClient(cfg config.InfluxDBConfig, logger *logger.Logger) (*InfluxClient, error) {
	client := influxdb2.NewClient(cfg.URL, cfg.Token)
	return &InfluxClient{
		client: client,
		org:    cfg.Org,
		bucket: cfg.Bucket,
		logger: logger,
	}, nil
}

func (c *InfluxClient) Save(data models.CollarMetrics) error {
	writeAPI := c.client.WriteAPIBlocking(c.org, c.bucket)

	petIDStr := strconv.FormatInt(data.PetID, 10)

	c.logger.Debug("Saving metrics to InfluxDB ", " pet_id: ", data.PetID, ", collar_id: ", data.CollarID)

	c.logger.Debug("pet_id: ", petIDStr, ", collar_id: ", data.CollarID,
		", temperature: ", data.Temperature,
		", heart_rate: ", data.HeartRate,
		", lat: ", data.Location.Lat,
		", lon: ", data.Location.Lon)

	point := influxdb2.NewPointWithMeasurement("pet_data").
		AddTag("pet_id", petIDStr).
		AddTag("collar_id", data.CollarID).
		AddField("temperature", data.Temperature).
		AddField("heart_rate", data.HeartRate).
		AddField("lat", data.Location.Lat).
		AddField("lon", data.Location.Lon).
		SetTime(data.Timestamp)

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		c.logger.Error("Failed to write to InfluxDB,", " error: ", err, ", pet_id: ", data.PetID, ", collar_id: ", data.CollarID)
		return fmt.Errorf("failed to write point to InfluxDB: %w", err)
	}

	c.logger.Info("Saved metrics to InfluxDB ", " pet_id: ", data.PetID, ", collar_id: ", data.CollarID)
	return nil
}

func (c *InfluxClient) GetMetrics(ctx context.Context, petID int64, duration time.Duration) ([]models.CollarMetrics, error) {
	queryAPI := c.client.QueryAPI(c.org)
	petIDStr := strconv.FormatInt(petID, 10)

	query := fmt.Sprintf(`
    from(bucket: "%s")
        |> range(start: -%s)
        |> filter(fn: (r) => r._measurement == "pet_data")
        |> filter(fn: (r) => r.pet_id == "%s")
        |> pivot(
            rowKey: ["_time"],
            columnKey: ["_field"],
            valueColumn: "_value"
        )
	`, c.bucket, duration, petIDStr)

	c.logger.Debug("Executing InfluxDB query", ", pet_id: ", petID, ", query: ", query)

	result, err := queryAPI.Query(ctx, query)
	if err != nil {
		c.logger.Error("InfluxDB query ", "error: ", err, ", pet_id: ", petID)
		return nil, fmt.Errorf("failed to query InfluxDB: %w", err)
	}
	defer result.Close()

	var metrics []models.CollarMetrics

	for result.Next() {
		record := result.Record()
		values := record.Values()
		c.logger.Debug("Processing InfluxDB record, ", "pet_id: ", petID, ", record_time: ", record.Time(), ", collar_id: ", record.ValueByKey("collar_id"))

		collarID, _ := values["collar_id"].(string)
		if collarID == "" {
			c.logger.Warn("Missing collar_id in record", "pet_id", petID, "time", record.Time())
			continue
		}
		metric := models.CollarMetrics{
			PetID:       petID,
			CollarID:    collarID,
			Timestamp:   record.Time(),
			Temperature: getFloat(values, "temperature"),
			HeartRate:   int(getFloat(values, "heart_rate")),
			Location: models.Location{
				Lat: getFloat(values, "lat"),
				Lon: getFloat(values, "lon"),
			},
		}

		metrics = append(metrics, metric)
	}

	if result.Err() != nil {
		c.logger.Error("InfluxDB result error", "error", result.Err(), "pet_id", petID)
		return nil, fmt.Errorf("failed to process InfluxDB result: %w", result.Err())
	}

	if len(metrics) == 0 {
		c.logger.Info("No metrics found for pet", "pet_id", petID)
	}

	c.logger.Debug("Returning metrics", "pet_id", petID, "count", len(metrics))
	return metrics, nil
}

func getFloat(values map[string]any, key string) float64 {
	if val, ok := values[key]; ok {
		switch v := val.(type) {
		case float64:
			return v
		case int64:
			return float64(v)
		case string:
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				return f
			}
		}
	}
	return 0
}

func (c *InfluxClient) Close() {
	c.client.Close()
}
