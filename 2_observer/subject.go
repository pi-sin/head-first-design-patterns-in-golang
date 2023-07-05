package main

// Subject interface
type Subject interface {
	/**
	* Both of these methods take an Observer as argument
	* that is, the Observer to be registered or removed.
	 */
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	/**
	* This method is called to notify all the observers
	* when the Subject's state has changed
	 */
	NotifyObservers()
}

// WeatherData struct
type WeatherData struct {
	observers                       []Observer
	temperature, humidity, pressure float64
}

// NewWeatherData creates a new WeatherData instance
func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

// RegisterObserver registers an observer
func (wd *WeatherData) RegisterObserver(observer Observer) {
	wd.observers = append(wd.observers, observer)
}

// RemoveObserver removes an observer
func (wd *WeatherData) RemoveObserver(observer Observer) {
	index := -1
	for i, obs := range wd.observers {
		if obs == observer {
			index = i
			break
		}
	}
	if index >= 0 {
		wd.observers = append(wd.observers[:index], wd.observers[index+1:]...)
	}
}

// NotifyObservers notifies all observers by calling their update() method
func (wd *WeatherData) NotifyObservers() {
	for _, observer := range wd.observers {
		observer.Update(wd.temperature, wd.humidity, wd.pressure)
	}
}

// We notify the Observers when measurements are updated
func (wd *WeatherData) measurementsChanged() {
	wd.NotifyObservers()
}

// SetMeasurements sets the weather measurements and notifies observers
func (wd *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.measurementsChanged()
}
