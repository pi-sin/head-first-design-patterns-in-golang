package main

import (
	"fmt"
)

type simplePizzaFactory struct{}

func (spf *simplePizzaFactory) createPizza(pizzaType string) (iPizza, error) {
	switch pizzaType {
	case "cheese":
		return newCheesePizza(), nil
	case "greek":
		return newGreekPizza(), nil
	case "pepperoni":
		return newPepperoniPizza(), nil
	}

	return nil, fmt.Errorf("Invalid pizza type")
}

type pizzaFactory interface {
	createPizza(pizzaType string) (iPizza, error)
}

type nyPizzaFactory struct {}

func (nypf *nyPizzaFactory) createPizza(pizzaType string) (iPizza, error) {
	switch pizzaType {
	case "cheese":
		return newCheesePizza(), nil
	case "greek":
		return newGreekPizza(), nil
	case "pepperoni":
		return newPepperoniPizza(), nil
	}

	return nil, fmt.Errorf("Invalid pizza type")
}