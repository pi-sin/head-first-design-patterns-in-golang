package main

import "fmt"

/**
* Observer interface is implemented by all observers,
* Here we're passing the measurements to the observers
* from the Subject when a weather measurement changes
 */
type Observer interface {
	Update(temp, humidity, pressure float64)
}

// CurrentConditionsDisplay struct
type CurrentConditionsDisplay struct {
	temperature, humidity float64
}

// NewCurrentConditionsDisplay creates a new CurrentConditionsDisplay instance
func NewCurrentConditionsDisplay() *CurrentConditionsDisplay {
	return &CurrentConditionsDisplay{}
}

// Update save the measurements and displays the required information
func (ccd *CurrentConditionsDisplay) Update(temp, humidity, pressure float64) {
	ccd.temperature = temp
	ccd.humidity = humidity
	ccd.Display()
}

// Display displays the current conditions
func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.2f°F temperature and %.2f%% humidity\n", ccd.temperature, ccd.humidity)
}

// StatisticsDisplay struct
type StatisticsDisplay struct {
	minTemp, maxTemp, avgTemp float64
	temperatureCount          int
}

// NewStatisticsDisplay creates a new StatisticsDisplay instance
func NewStatisticsDisplay() *StatisticsDisplay {
	return &StatisticsDisplay{}
}

// Update calculates avg, min and max temp and displays the information
func (sd *StatisticsDisplay) Update(temp, humidity, pressure float64) {
	if sd.temperatureCount == 0 {
		sd.minTemp = temp
		sd.maxTemp = temp
		sd.avgTemp = temp
	} else {
		if temp < sd.minTemp {
			sd.minTemp = temp
		}
		if temp > sd.maxTemp {
			sd.maxTemp = temp
		}
		sd.avgTemp = (sd.avgTemp*float64(sd.temperatureCount) + temp) / float64(sd.temperatureCount+1)
	}

	sd.temperatureCount++
	sd.Display()
}

// Display displays the statistics
func (sd *StatisticsDisplay) Display() {
	fmt.Printf("Temperature statistics: Min %.2f°F, Max %.2f°F, Average %.2f°F\n", sd.minTemp, sd.maxTemp, sd.avgTemp)
}
