package main

import (
	"mqtt-golang-subscriber/adapter"
	"mqtt-golang-subscriber/controllers"
	"mqtt-golang-subscriber/db"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	mqttConnection := adapter.NewConnection(os.Getenv(adapter.MqttClientName))
	mqttConnection.Subscribe(os.Getenv(adapter.MqttTopicName))

	influxDBConnection := db.NewConnection()
	//influxDBConnection.Insert()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", controllers.HealthControllerHandler(mqttConnection, influxDBConnection))
	}

	//influxDBConnection := db.NewConnection()
	//influxDBConnection.Insert()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "Not found"})
	})

	router.Run(":8080")
}
