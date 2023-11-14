package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var WeatherAPIKey = os.Getenv("WEATHER_API_KEY")

func LookupWeather(input string, forecast bool) (string, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=3&aqi=no&alerts=no", WeatherAPIKey, input), nil)
	if err != nil {
		log.Print("Error building request")
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("Error sending request")
		return "", err
	}

	var respJSON Weather
	json.NewDecoder(resp.Body).Decode(&respJSON)

	currentWeather := fmt.Sprintf(
		"In %s, %s it is %s°F and %s.",
		respJSON.Location.City,
		respJSON.Location.State,
		respJSON.Current.Temperature,
		respJSON.Current.Condition.Type,
	)

	day1Forecast := fmt.Sprintf(
		"On %s, it will be %s with a high of %s°F and a low of %s°F.",
		respJSON.Forecast.ForecastDay[0].Date,
		respJSON.Forecast.ForecastDay[0].Day.Condition.Type,
		respJSON.Forecast.ForecastDay[0].Day.HighTemp,
		respJSON.Forecast.ForecastDay[0].Day.LowTemp,
	)

	day2Forecast := fmt.Sprintf(
		"On %s, it will be %s with a high of %s°F and a low of %s°F.",
		respJSON.Forecast.ForecastDay[1].Date,
		respJSON.Forecast.ForecastDay[1].Day.Condition.Type,
		respJSON.Forecast.ForecastDay[1].Day.HighTemp,
		respJSON.Forecast.ForecastDay[1].Day.LowTemp,
	)

	day3Forecast := fmt.Sprintf(
		"On %s, it will be %s with a high of %s°F and a low of %s°F.",
		respJSON.Forecast.ForecastDay[2].Date,
		respJSON.Forecast.ForecastDay[2].Day.Condition.Type,
		respJSON.Forecast.ForecastDay[2].Day.HighTemp,
		respJSON.Forecast.ForecastDay[2].Day.LowTemp,
	)

	forecastedWeather := fmt.Sprintf("%s\n\nHere is the forecast for the next three days:\n\n%s\n%s\n%s", currentWeather, day1Forecast, day2Forecast, day3Forecast)

	if !forecast {
		return currentWeather, nil
	} else {
		return forecastedWeather, nil
	}
}
