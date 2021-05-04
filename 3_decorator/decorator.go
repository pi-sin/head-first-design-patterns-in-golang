package main

type mocha struct {
	beverage beverage
}

func (m *mocha) description() string {
	return m.beverage.description() + ", Mocha"
}

func (m *mocha) cost() float32 {
	return m.beverage.cost() + .2
}

type milk struct {
	beverage beverage
}

func (m *milk) description() string {
	return m.beverage.description() + ", Milk"
}

func (m *milk) cost() float32 {
	return m.beverage.cost() + .1
}

type soy struct {
	beverage beverage
}

func (s *soy) description() string {
	return s.beverage.description() + ", Soy"
}

func (s *soy) cost() float32 {
	return s.beverage.cost() + .15
}

type whip struct {
	beverage beverage
}

func (w *whip) description() string {
	return w.beverage.description() + ", Whip"
}

func (w *whip) cost() float32 {
	return w.beverage.cost() + .1
}
