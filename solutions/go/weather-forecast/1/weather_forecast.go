// Package weather provides tools to forecast the current weather condition for a given city.
package weather

// CurrentCondition represents the current weather condition as a string.
var CurrentCondition string
// CurrentLocation represents the current location as a string.
var CurrentLocation string

// Forecast takes a city and a condition string parameters and returns a string value describing the current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
