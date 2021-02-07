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

	influxDBConnection := db.NewConnection()

	mqttConnection := adapter.NewConnection(os.Getenv(adapter.MqttClientName))
	mqttConnection.Subscribe(influxDBConnection, os.Getenv(adapter.MqttTopicName))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", controllers.HealthControllerHandler(mqttConnection, influxDBConnection))
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "Not found"})
	})

	router.Run(":8080")
}
