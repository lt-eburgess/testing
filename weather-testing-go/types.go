package weather

import "encoding/json"

type Condition struct {
	Type string `json:"text"`
}

type CurrentConditions struct {
	Temperature json.Number `json:"temp_f"`
	Condition   Condition   `json:"condition"`
}

type Day struct {
	HighTemp  json.Number `json:"maxtemp_f"`
	LowTemp   json.Number `json:"mintemp_f"`
	Condition Condition   `json:"condition"`
}

type ForecastDay []struct {
	Date string `json:"date"`
	Day  Day    `json:"day"`
}

type Forecast struct {
	ForecastDay ForecastDay `json:"forecastday"`
}

type Location struct {
	City  string `json:"name"`
	State string `json:"region"`
}

type Weather struct {
	Location Location          `json:"location"`
	Current  CurrentConditions `json:"current"`
	Forecast Forecast          `json:"forecast"`
}
