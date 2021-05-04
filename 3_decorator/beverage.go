package main

/**
 * For our case, we can use interface for beverage
 * instead of an abstract class.
 */
type beverage interface {
	description() string
	cost() float32
}

type espresso struct{}

func (e *espresso) description() string {
	return "Espresso"
}

func (e *espresso) cost() float32 {
	return 1.99
}

type houseBlend struct{}

func (hb *houseBlend) description() string {
	return "House Blend Coffee"
}

func (hb *houseBlend) cost() float32 {
	return 0.89
}

type darkRoast struct{}

func (dr *darkRoast) description() string {
	return "Dark Roast Coffee"
}

func (dr *darkRoast) cost() float32 {
	return 0.99
}

type decaf struct{}

func (d *decaf) description() string {
	return "Decaf"
}

func (d *decaf) cost() float32 {
	return 1.05
}
