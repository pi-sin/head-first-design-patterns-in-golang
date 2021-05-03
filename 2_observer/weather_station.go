package main

func main() {
	weatherData := newWeatherData()

	currentConditionDisplay := newCurrentConditionDisplay()
	weatherData.registerObserver(currentConditionDisplay)

	statisticsDisplay := newStatisticsDisplay()
	weatherData.registerObserver(statisticsDisplay)
	//forecastDisplay := newForecastDisplay()
	//weatherData.registerObserver(forecastDisplay)

	weatherData.setMeasurements(80, 65, 30.4)
	weatherData.setMeasurements(82, 70, 29.2)
	weatherData.setMeasurements(77, 90, 29.2)

	weatherData.deRegisterObserver(statisticsDisplay)

	weatherData.setMeasurements(84, 80, 32.6)

	weatherData.registerObserver(statisticsDisplay)

	weatherData.setMeasurements(84, 80, 32.6)

}
