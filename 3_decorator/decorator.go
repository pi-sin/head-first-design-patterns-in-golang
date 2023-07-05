package main

// Decorator
type CondimentDecorator interface {
	Beverage
	// New method added to the CondimentDecorator interface
	AddOnDescription() string
}

// ConcreteDecorator
type Milk struct {
	beverage Beverage
}

func NewMilk(beverage Beverage) Beverage {
	return &Milk{beverage: beverage}
}

func (m *Milk) Cost() float64 {
	return 0.1 + m.beverage.Cost()
}

func (m *Milk) GetDescription() string {
	return m.beverage.GetDescription() + m.AddOnDescription()
}

func (m *Milk) AddOnDescription() string {
	return "\n +AddOn Milk"
}

// ConcreteDecorator
type Mocha struct {
	beverage Beverage
}

func NewMocha(beverage Beverage) Beverage {
	return &Mocha{beverage: beverage}
}

func (m *Mocha) Cost() float64 {
	return 0.2 + m.beverage.Cost()
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + m.AddOnDescription()
}

func (m *Mocha) AddOnDescription() string {
	return "\n +AddOn Mocha"
}

// ConcreteDecorator
type Soy struct {
	beverage Beverage
}

func NewSoy(beverage Beverage) Beverage {
	return &Soy{beverage: beverage}
}

func (s *Soy) Cost() float64 {
	return 0.15 + s.beverage.Cost()
}

func (s *Soy) GetDescription() string {
	return s.beverage.GetDescription() + s.AddOnDescription()
}

func (s *Soy) AddOnDescription() string {
	return "\n +AddOn Soy"
}
