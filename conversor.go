package main

import "fmt"

const BOLING_WATER_TEMPERATURE_IN_KELVIN = 373.2

func main() {
	fmt.Println("\nThe boiling point of water in celcius is:", convertTemperature(BOLING_WATER_TEMPERATURE_IN_KELVIN))
}

func convertTemperature(degreesKelvin float32) int16 {

	var convertedTemperature int16 = int16(degreesKelvin - 273)

	return convertedTemperature
}
