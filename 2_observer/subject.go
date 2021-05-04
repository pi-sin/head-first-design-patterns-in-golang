package main

type subject interface {
	/**
	 * Both of these methods take an Observer as argument
	 */
	registerObserver(observer)
	deregisterObserver(observer)
	/**
	 * This method is called to notify all the observers
	 * when the Subject's state has changed
	 */
	notifyObservers()
}

type weatherData struct {
	/**
	 * Implementation of set in Go
	 * Instantiating in the constructor (newWeatherData)
	 */
	observers   map[observer]bool
	temperature float32
	humidity    float32
	pressure    float32
}

func newWeatherData() *weatherData {
	return &weatherData{
		observers: make(map[observer]bool),
	}
}

/**
 * WeatherData now implements the Subject interface
 */
func (w *weatherData) registerObserver(o observer) {
	/**
	 * When an observer registers, we add it in
	 * the map with value set as true
	 */
	w.observers[o] = true
}

func (w *weatherData) deregisterObserver(o observer) {
	/**
	 * When an observer deregisters, we remove it
	 * from the map after checking if the value exists
	 */
	if _, ok := w.observers[o]; ok {
		delete(w.observers, o)
	}
}

func (w *weatherData) notifyObservers() {
	/**
	 * This is where we tell all the observers about the state
	 * by calling the common update method
	 */
	for observer := range w.observers {
		observer.update(w.temperature, w.humidity, w.pressure)
	}
}

/**
 * We notify the Observers when we get updated
 * measurements from the Weather Station
 */
func (w *weatherData) measurementsChanged() {
	w.notifyObservers()
}

/**
 * Dummy method to test our display elements.
 * Setting the measurements not via a device.
 */
func (w *weatherData) setMeasurements(temp float32, humidity float32, pressure float32) {
	w.temperature = temp
	w.humidity = humidity
	w.pressure = pressure
	w.measurementsChanged()
}
