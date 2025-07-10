package mqvizclient

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Default handler if none provided
var DefaultMessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received on topic %s: %s\n", msg.Topic(), string(msg.Payload()))
}

var DefaultConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected to MQTT broker.")
}
