package main

import (
	"fmt"
	"log"
	"weather"
	"zipcodes"
)

func main() {

	input := "94088"
	forecast := true

	valid, err := zipcodes.ZipcodeIsValid(input)
	if !valid && err != nil {
		log.Print(err)
	} else {
		userResponse, err := weather.LookupWeather(input, forecast)
		if err != nil {
			log.Print(err)
		}

		fmt.Print(userResponse)
	}
}
