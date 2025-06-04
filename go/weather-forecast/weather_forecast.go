// Package weather provides tools for searching 
// for weather forecasts in your region.
package weather

// CurrentCondition string represents the current condition.
var CurrentCondition string
// CurrentLocation string represents the current location.
var CurrentLocation string

// Forecast returns a string containing your location and the current weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
