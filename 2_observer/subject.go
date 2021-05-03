package main

type subject interface {
	registerObserver(observer)
	deRegisterObserver(observer)
	notifyObservers()
}

type weatherData struct {
	observers   map[observer]bool
	temperature float32
	humidity    float32
	pressure    float32
}

func newWeatherData() *weatherData  {
	return &weatherData{
		observers: make(map[observer]bool),
	}
}

func (w *weatherData) registerObserver(o observer) {
	w.observers[o] = true
}

func (w *weatherData) deRegisterObserver(o observer) {
	if _, ok := w.observers[o]; ok {
		delete(w.observers, o)
	}
}

func (w *weatherData) notifyObservers() {
	for observer := range w.observers {
		observer.update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *weatherData) setMeasurements(temp float32, humidity float32, pressure float32) {
	w.temperature = temp
	w.humidity = humidity
	w.pressure = pressure
	w.notifyObservers()
}