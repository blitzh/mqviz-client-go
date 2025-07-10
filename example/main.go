package example

import (
	"fmt"
	"mqvizclient"
)

func main() {
	err := mqvizclient.Connect(mqvizclient.Config{
		Broker:    "tcp://localhost:1883",
		ClientID:  "my-dashboard-client",
		Username:  "",
		Password:  "",
		OnConnect: mqvizclient.DefaultConnectHandler,
		OnMessage: mqvizclient.DefaultMessageHandler,
	})
	if err != nil {
		panic(err)
	}

	err = ListenDashboardResponse("my-dashboard-client", func(resp DashboardResponse) {
		fmt.Println("Dashboard Data Received:")
		fmt.Printf("Status: %s\n", resp.Status)
		fmt.Printf("Message: %s\n", resp.Message)
		fmt.Printf("Data: %+v\n", resp.Data)
	})

	if err != nil {
		panic(err)
	}

	select {} // keep alive
}
