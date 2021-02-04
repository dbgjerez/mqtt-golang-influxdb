package adapter

import (
	"fmt"
	"log"
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
	opts.SetDefaultPublishHandler(messagePubHandler)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln("Connect problem: ", token.Error())
	}
	conn = &MqttConnection{client}
	return conn
}

func (conn *MqttConnection) Subscribe(topic string) {
	token := conn.mqttClient.Subscribe(topic, 1, nil)
	token.Wait()
	log.Println("Subscribed to topic: ", topic)
}

func (con *MqttConnection) IsConnected() bool {
	return con.mqttClient.IsConnected()
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s", msg.Payload(), msg.Topic())
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Println("Connection lost: ", err)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Mqtt connected")
}
