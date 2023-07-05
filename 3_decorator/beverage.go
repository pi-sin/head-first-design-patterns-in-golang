package main

// Component
type Beverage interface {
	Cost() float64
	GetDescription() string
}

// ConcreteComponent
type Espresso struct{}

func NewEspresso() Beverage {
	return &Espresso{}
}

func (e *Espresso) Cost() float64 {
	return 1.99
}

func (e *Espresso) GetDescription() string {
	return "Espresso"
}

// ConcreteComponent: DarkRoast
type DarkRoast struct{}

func NewDarkRoast() Beverage {
	return &DarkRoast{}
}

func (d *DarkRoast) Cost() float64 {
	return 0.99
}

func (d *DarkRoast) GetDescription() string {
	return "Dark Roast Coffee"
}
