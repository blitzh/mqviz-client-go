package example

type DashboardResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` // atau pakai struct detail jika spesifik
}
