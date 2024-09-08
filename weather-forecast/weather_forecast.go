// Package weather check the weather in a city.
package weather

// CurrentCondition is the actual condition.
var CurrentCondition string

// CurrentLocation contains the location.
var CurrentLocation string

// Forecast return the actual condition in a determined city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
