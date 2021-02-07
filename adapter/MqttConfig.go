package adapter

import (
	"encoding/json"
	"fmt"
	"log"
	"mqtt-golang-subscriber/db"
	"mqtt-golang-subscriber/models"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// mqtt constants
const (
	BrokerHostFormat = "tcp://%s:%s" // host yo mqtt
	MqttHost         = "MQTT_HOST"   // host env variable
	MqttPort         = "MQTT_PORT"
	MqttClientName   = "MQTT_CLIENT_NAME"
	MqttTopicName    = "MQTT_TOPIC_NAME"
)

var host = os.Getenv(MqttHost)
var port = os.Getenv(MqttPort)

type MqttConnection struct {
	mqttClient mqtt.Client
}

func NewConnection(clientId string) (conn *MqttConnection) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf(BrokerHostFormat, host, port))
	opts.SetClientID(clientId)
	opts.AutoReconnect = true
	opts.OnConnectionLost = connectLostHandler
	opts.OnConnect = connectHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln("Connect problem: ", token.Error())
	}
	conn = &MqttConnection{client}
	return conn
}

func (conn *MqttConnection) Subscribe(influxConn *db.InfluxDBConnection, topic string) {
	token := conn.mqttClient.Subscribe(topic, 1, onMessageReceived(influxConn))
	token.Wait()
	log.Println("Subscribed to topic: ", topic)
}

func (con *MqttConnection) IsConnected() bool {
	connected := con.mqttClient.IsConnected()
	if !connected {
		log.Println("Healthcheck MQTT fails")
	}
	return connected
}

func onMessageReceived(influxConn *db.InfluxDBConnection) func(client mqtt.Client, msg mqtt.Message) {
	return func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message: %s from topic: %s", msg.Payload(), msg.Topic())

		event := models.ChipEvent{}

		err := json.Unmarshal([]byte(msg.Payload()), &event)
		if err != nil {
			log.Println("Unmarshal message fails: ", err)
		}

		influxConn.Insert(&event)
	}
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Println("Connection lost: ", err)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Mqtt connected")
}
