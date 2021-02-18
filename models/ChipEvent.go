package models

type ChipEvent struct {
	Chip    string `json:"chip"`
	Battery string `json:"battery"`
	Sensors []struct {
		Sensor   string `json:"sensor"`
		Time     int64  `json:"time"`
		Humidity int    `json:"humidity"`
	} `json:"sensors"`
}
