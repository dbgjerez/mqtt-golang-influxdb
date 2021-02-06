package models

// health status
const (
	HealhStatusUp   = "UP"
	HealhStatusDown = "DOWN"
)

type Health struct {
	Status string `json:"status"`
}
