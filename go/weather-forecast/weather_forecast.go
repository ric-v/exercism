// Package weather forecast.
package weather

// CurrentCondition - describes current weather condition.
var CurrentCondition string

// CurrentLocation - describes current location.
var CurrentLocation string

// Forecast returns a weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
