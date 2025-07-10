package example

import (
	"encoding/json"
	"fmt"
	"mqvizclient"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ListenDashboardResponse(clientID string, handler func(DashboardResponse)) error {
	topic := DashboardResponseTopic(clientID)
	return mqvizclient.Subscribe(topic, func(c mqtt.Client, m mqtt.Message) {
		var resp DashboardResponse
		err := json.Unmarshal(m.Payload(), &resp)
		if err != nil {
			fmt.Println("Failed to parse dashboard response:", err)
			return
		}
		handler(resp)
	})
}
