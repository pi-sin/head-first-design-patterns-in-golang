package main

func main() {
	weatherData := NewWeatherData()

	currentConditionsDisplay := NewCurrentConditionsDisplay()
	statisticsDisplay := NewStatisticsDisplay()

	weatherData.RegisterObserver(currentConditionsDisplay)
	weatherData.RegisterObserver(statisticsDisplay)

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.2)
}
