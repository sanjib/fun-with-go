package main

import (
	"encoding/json"
	"fmt"
)

type WeatherData struct {
	LocationName string `json:"location"`
	Weather      string `json:"weather"`
	Temperature  int    `json:"temp"`
	Celsius      bool   `json:"celsius,omitempty"`
	TempForecast []int  `json:"temp_forecast"`
}

func main() {
	w1 := WeatherData{
		LocationName: "Cobbs Village",
		Weather:      "sunny",
		Temperature:  80,
		Celsius:      false,
		TempForecast: []int{27, 25, 28},
	}
	data, err := json.MarshalIndent(w1, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
