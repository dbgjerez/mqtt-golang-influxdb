package controllers

import (
	"mqtt-golang-subscriber/adapter"
	"mqtt-golang-subscriber/db"
	"mqtt-golang-subscriber/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func HealthControllerHandler(mqttConn *adapter.MqttConnection, influxConn *db.InfluxDBConnection) func(c *gin.Context) {
	return func(c *gin.Context) {
		h := models.Health{}
		if mqttConn.IsConnected() && influxConn.IsConnected() {
			h.Status = models.HealhStatusUp
		} else {
			h.Status = models.HealhStatusUp
		}
		c.JSON(http.StatusOK, h)
	}
}
