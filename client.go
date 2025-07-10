package mqvizclient

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client

type Config struct {
	Broker    string
	ClientID  string
	Username  string
	Password  string
	OnConnect mqtt.OnConnectHandler
	OnMessage mqtt.MessageHandler
}

func Connect(cfg Config) error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.Broker)
	opts.SetClientID(cfg.ClientID)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.OnConnect = cfg.OnConnect
	opts.DefaultPublishHandler = cfg.OnMessage
	opts.ConnectTimeout = 5 * time.Second

	client = mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("MQTT connection error: %w", token.Error())
	}
	return nil
}

func Publish(topic string, payload []byte) error {
	if client == nil || !client.IsConnected() {
		return fmt.Errorf("client not connected")
	}
	token := client.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}

func Subscribe(topic string, callback mqtt.MessageHandler) error {
	if client == nil || !client.IsConnected() {
		return fmt.Errorf("client not connected")
	}
	token := client.Subscribe(topic, 0, callback)
	token.Wait()
	return token.Error()
}
