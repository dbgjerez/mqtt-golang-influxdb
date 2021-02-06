package controllers

import (
	"mqtt-golang-subscriber/adapter"
	"mqtt-golang-subscriber/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

type Health struct {
	Status string `json:"status"`
}

func HealthControllerHandler(mqttConn *adapter.MqttConnection, influxConn *db.InfluxDBConnection) func(c *gin.Context) {
	return func(c *gin.Context) {
		h := Health{}
		if mqttConn.IsConnected() && influxConn.IsConnected() {
			h.Status = "UP"
		} else {
			h.Status = "DOWN"
		}
		c.JSON(http.StatusOK, h)
	}
}

func (h *HealthController) Default(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}
