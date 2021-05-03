package main

func main() {
	/**
	 * Creating the weather object
	 */
	weatherData := newWeatherData()

	/**
	 * Creating the displays and registering as
	 * observer to the Subject
	 */
	currentConditionDisplay := newCurrentConditionDisplay()
	weatherData.registerObserver(currentConditionDisplay)

	statisticsDisplay := newStatisticsDisplay()
	weatherData.registerObserver(statisticsDisplay)
	//forecastDisplay := newForecastDisplay()
	//weatherData.registerObserver(forecastDisplay)

	/**
	 * Simulate new weather measurements
	 */
	weatherData.setMeasurements(80, 65, 30.4)
	weatherData.setMeasurements(82, 70, 29.2)
	weatherData.setMeasurements(77, 90, 29.2)

	/**
	 * Deregistering the statisticsDisplay
	 * Change in measurements won't notify this observer
	 */
	weatherData.deregisterObserver(statisticsDisplay)

	weatherData.setMeasurements(84, 80, 32.6)

	weatherData.registerObserver(statisticsDisplay)

	weatherData.setMeasurements(84, 80, 32.6)
}
