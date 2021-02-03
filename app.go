package main

import (
	"mqtt-golang-subscriber/adapter"
	"mqtt-golang-subscriber/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		health := new(controllers.HealthController)
		v1.GET("/health", health.Default)
	}

	conn := adapter.NewConnection(os.Getenv(adapter.MqttClientName))
	conn.Subscribe(os.Getenv(adapter.MqttTopicName))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "Not found"})
	})

	router.Run(":8080")
}
