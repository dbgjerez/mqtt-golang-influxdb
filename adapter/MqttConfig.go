package adapter

import (
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// mqtt constants
const (
	BrokerHost = "tcp://%s:%s" // host yo mqtt
)

var host = os.Getenv("MQTT_HOST")
var port = os.Getenv("MQTT_PORT")

type MqttConnection struct {
	mqttClient mqtt.Client
}

func NewConnection(clientId string) (conn *MqttConnection) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf(host, host, port))
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

func (con *MqttConnection) Subscribe(topic string) {
	token := con.mqttClient.Subscribe(topic, 1, nil)
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
