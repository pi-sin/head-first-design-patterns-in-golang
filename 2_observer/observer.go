package main

import "fmt"

/**
 * Observer interface is implemented by all observers,
 * Here we're passing the measurements to the observers
 * from the Subject when a weather measurement changes
 */
type observer interface {
	update(temp float32, humidity float32, pressure float32)
}

type currentConditionDisplay struct {
	temperature float32
	humidity    float32
	pressure    float32
}

func newCurrentConditionDisplay() *currentConditionDisplay {
	return &currentConditionDisplay{}
}

/**
 * This implements Observer and DisplayElement interfaces to get changes from WeatherData
 * and show the information based on its functionality respectively
 */
func (ccd *currentConditionDisplay) update(temp float32, humidity float32, pressure float32) {
	ccd.temperature = temp
	ccd.humidity = humidity
	ccd.pressure = pressure
	/**
	 * When update() is called, we save the measurements
	 * and call display() to show the required information
	 */
	ccd.display()
}

func (ccd *currentConditionDisplay) display() {
	fmt.Printf("Current conditions: Temperature:%.2f, Humidity:%.2f and Pressure:%.2f\n", ccd.temperature, ccd.humidity, ccd.pressure)
}

type statisticsDisplay struct {
	count   uint32
	avgTemp float32
	maxTemp float32
	minTemp float32
}

func newStatisticsDisplay() *statisticsDisplay {
	return &statisticsDisplay{}
}

/**
 * This implements the avg, max and min temperatures.
 * When update is called, it does the calculation and
 * call the display() method to show the required information
 */
func (sd *statisticsDisplay) update(temp float32, humidity float32, pressure float32) {
	sd.count++
	sd.avgTemp -= (sd.avgTemp - temp) / float32(sd.count)

	if sd.maxTemp < temp || sd.maxTemp == 0.0 {
		sd.maxTemp = temp
	}

	if sd.minTemp > temp || sd.minTemp == 0.0 {
		sd.minTemp = temp
	}

	sd.display()
}

func (sd *statisticsDisplay) display() {
	fmt.Printf("Avg/Max/Min Temperature:%.2f, %.2f, %.2f\n", sd.avgTemp, sd.maxTemp, sd.minTemp)
}
