package models

type ChipEvent struct {
	Chip    string `json:"chip"`
	Battery int64  `json:"battery,string,omitempty"`
	Sensors []struct {
		Sensor   string `json:"sensor"`
		Time     int64  `json:"time"`
		Humidity int    `json:"humidity"`
	} `json:"sensors"`
}
