package models

import "time"

//ChartData ...
type EventData struct {
	Date  []time.Time
	Value []float64
}

//ChartData ...
type EventDataNice struct {
	Date  []string
	Value []float64
}

type EventGrid struct {
	DeviceID  int32
	Coord     string
	Temp      float32
	Humidity  float32
	WindSpeed float32
	Date      time.Time
}

type EventNiceGrid struct {
	DeviceID  int32   `json:"deviceID"`
	Coord     string  `json:"coord"`
	Temp      float32 `json:"temp"`
	Humidity  float32 `json:"humidity"`
	WindSpeed float32 `json:"windSpeed"`
	Date      string  `json:"date"`
}
