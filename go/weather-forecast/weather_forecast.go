// Package weather: help forecast weather in Goblinocus.
package weather

// CurrentCondition variable that tracks the current condition
// of a city in Goblinocus.
var CurrentCondition string
// CurrentLocation The location of place or city in Goblinocus.
var CurrentLocation string

// Forecast tells you the weather of a city or location in Goblinocus.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
