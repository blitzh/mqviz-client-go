package example

import "fmt"

func DashboardResponseTopic(clientID string) string {
	return fmt.Sprintf("mqviz/%s/dashboard/response", clientID)
}
