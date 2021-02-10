package db

import (
	"context"
	"log"
	"mqtt-golang-subscriber/models"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// Influxdb constants
const (
	InfluxDBHost         = "INFLUXDB_HOST"
	InfluxDBDatabaseName = "INFLUXDB_DATABASE_NAME"
	InfluxDBMeasurement  = "INFLUXDB_MEASUREMENT"
)

type InfluxDBConnection struct {
	influxdbClient influxdb2.Client
}

func NewConnection() (conn *InfluxDBConnection) {
	influxdb2.DefaultOptions().HTTPClient()
	client := influxdb2.NewClient(os.Getenv(InfluxDBHost), "")
	conn = &InfluxDBConnection{client}
	return conn
}

func (conn *InfluxDBConnection) IsConnected() bool {
	_, err := conn.influxdbClient.Health(context.Background())
	if err != nil {
		log.Println("Healthcheck InfluxDB fails: ", err)
		return false
	}
	return true
}

func (conn *InfluxDBConnection) Insert(event *models.ChipEvent) {
	for _, elem := range event.Sensors {
		p := influxdb2.NewPointWithMeasurement(os.Getenv(InfluxDBMeasurement)).
			AddTag("chip", event.Chip).
			AddTag("sensor", elem.Sensor).
			AddField("humidity", elem.Humidity).
			SetTime(time.Unix(elem.Time, 0))

		writeAPI := conn.influxdbClient.WriteAPIBlocking("", os.Getenv(InfluxDBDatabaseName))
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			log.Println("Influxdb fails insert: ", err)
		}
	}
}
