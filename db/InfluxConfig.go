package db

import (
	"context"
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
	//influxdb2.NewPointWithMeasurement()
	conn = &InfluxDBConnection{client}
	return conn
}

func (conn *InfluxDBConnection) Insert() {
	p := influxdb2.NewPointWithMeasurement(InfluxDBMeasurement).
		AddTag("chip", "asdasd").
		AddTag("sensor", "fc28").
		AddField("humidity", 1023).
		SetTime(time.Now())

	writeAPI := conn.influxdbClient.WriteAPIBlocking("", os.Getenv(InfluxDBDatabaseName))
	writeAPI.WritePoint(context.Background(), p)
}
