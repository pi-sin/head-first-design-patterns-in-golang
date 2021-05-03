package main

import "fmt"

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

func (ccd *currentConditionDisplay) update(temp float32, humidity float32, pressure float32) {
	ccd.temperature = temp
	ccd.humidity = humidity
	ccd.pressure = pressure
	ccd.display()
}

func (ccd *currentConditionDisplay) display()  {
	fmt.Printf("Current conditions: Temperature:%.2f, Humidity:%.2f and Pressure:%.2f\n", ccd.temperature, ccd.humidity, ccd.pressure)
}

type statisticsDisplay struct {
	count uint32
	totalTemp float32
	maxTemp float32
	minTemp float32
}

func newStatisticsDisplay() *statisticsDisplay {
	return &statisticsDisplay{}
}

func (sd *statisticsDisplay) update(temp float32, humidity float32, pressure float32) {
	sd.count++
	sd.totalTemp += temp

	if sd.maxTemp < temp || sd.maxTemp == 0.0 {
		sd.maxTemp = temp
	}

	if sd.minTemp > temp || sd.minTemp == 0.0 {
		sd.minTemp = temp
	}

	sd.display()
}

func (sd *statisticsDisplay) display()  {
	fmt.Printf("Avg/Max/Min Temperature:%.2f, %.2f, %.2f\n", sd.totalTemp/float32(sd.count), sd.maxTemp, sd.minTemp)
}
