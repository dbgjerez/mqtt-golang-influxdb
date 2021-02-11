# MQTT client for iot

# Configuration
| Variable | Default value | Description |
| ------ | ------ | ------ |
| PORT | 8080 | Server port |
| GIN_MODE | debug | Gin gonic mode. (release for production mode) |
| MQTT_HOST | mqtt.server.com | Mqtt host |
| MQTT_PORT | 1883 | Mqtt port |
| MQTT_CLIENT_NAME | "" | Name of the ms when connect to Mqtt broker |
| MQTT_TOPIC_NAME | "" | Topic to suscription |
| INFLUXDB_HOST | "" | Influxdb host |
| INFLUXDB_DATABASE_NAME | "" | Influxdb database name |
| INFLUXDB_MEASUREMENT | "" | Influxdb measurement name for this nanoservice |

# Libraries
* Gin Gonic: Golang Framework
* InfluxDB client: Connection with InfluxDB https://github.com/influxdata/influxdb-client-go
* Mqtt client: Mqtt connection and listener https://github.com/eclipse/paho.mqtt.golang
* GoDotEnv: Library for env variables
