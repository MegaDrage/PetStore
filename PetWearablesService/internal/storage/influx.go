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

	c.logger.Debug("Saving metrics to InfluxDB ", " pet_id: ", data.PetID)

	c.logger.Debug("pet_id: ", data.PetID,
		", temperature: ", data.Temperature,
		", heart_rate: ", data.HeartRate,
		", lat: ", data.Location.Lat,
		", lon: ", data.Location.Lon)

	point := influxdb2.NewPointWithMeasurement("pet_data").
		AddTag("pet_id", data.PetID).
		AddField("temperature", data.Temperature).
		AddField("heart_rate", data.HeartRate).
		AddField("lat", data.Location.Lat).
		AddField("lon", data.Location.Lon).
		SetTime(data.Timestamp)

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		c.logger.Error("Failed to write to InfluxDB,", " error: ", err, ", pet_id: ", data.PetID,)
		return fmt.Errorf("failed to write point to InfluxDB: %w", err)
	}

	c.logger.Info("Saved metrics to InfluxDB", " pet_id: ", data.PetID)
	return nil
}

func (c *InfluxClient) GetMetrics(ctx context.Context, petID string, duration time.Duration) ([]models.CollarMetrics, error) {
    queryAPI := c.client.QueryAPI(c.org)
	
	basic_duration := 15 * time.Minute	

	if duration != 0 {
		basic_duration = duration
	}

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
    `, c.bucket, basic_duration, petID)

    c.logger.Debug("Executing InfluxDB query", "pet_id", petID, "query", query)

    result, err := queryAPI.Query(ctx, query)
    if err != nil {
        c.logger.Error("InfluxDB query error", "error", err, "pet_id", petID)
        return nil, fmt.Errorf("failed to query InfluxDB: %w", err)
    }
    defer result.Close()

    var tempMetrics []models.CollarMetrics
    for result.Next() {
        record := result.Record()
        values := record.Values()

        metric := models.CollarMetrics{
            PetID:       petID,
            Timestamp:   record.Time(),
            Temperature: getFloat(values, "temperature"),
            HeartRate:   int(getFloat(values, "heart_rate")),
            Location: models.Location{
                Lat: getFloat(values, "lat"),
                Lon: getFloat(values, "lon"),
            },
        }
        tempMetrics = append(tempMetrics, metric)
    }

    if result.Err() != nil {
        c.logger.Error("InfluxDB result error: ", result.Err(), ", pet_id: ", petID)
        return nil, fmt.Errorf("failed to process InfluxDB result: %w", result.Err())
    }

    if len(tempMetrics) == 0 {
        c.logger.Info("No metrics found for pet, ", "pet_id: ", petID)
        return nil, nil
    }

    c.logger.Debug("Collected metrics,", " pet_id: ", petID, ", count: ", len(tempMetrics))

    if duration == 0 {
        lastMetric := tempMetrics[len(tempMetrics)-1]
        c.logger.Info("Returning last metric, ", "pet_id: ", petID, ", timestamp: ", lastMetric.Timestamp)
        return []models.CollarMetrics{lastMetric}, nil
    }

    var sumTemp, sumHeartRate, sumLat, sumLon float64
    count := float64(len(tempMetrics))
    for _, m := range tempMetrics {
        sumTemp += m.Temperature
        sumHeartRate += float64(m.HeartRate)
        sumLat += m.Location.Lat
        sumLon += m.Location.Lon
    }

    avgMetric := models.CollarMetrics{
        PetID:       petID,
        Timestamp:   time.Now(),
        Temperature: sumTemp / count,
        HeartRate:   int(sumHeartRate/count + 0.5),
        Location: models.Location{
            Lat: sumLat / count,
            Lon: sumLon / count,
        },
    }

    c.logger.Info("Returning averaged metrics, ", "pet_id: ", petID, ", avg_temperature: ", avgMetric.Temperature, ", avg_heart_rate:", avgMetric.HeartRate)
    return []models.CollarMetrics{avgMetric}, nil
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
