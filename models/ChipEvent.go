package models

type ChipEvent struct {
	Chip    string `json:"chip"`
	Sensors []struct {
		Sensor   string `json:"sensor"`
		Time     int    `json:"time"`
		Humidity int    `json:"humidity"`
	} `json:"sensors"`
}
