package main

import "fmt"

// Abstract Interface
type iPizzaStore interface {
	orderPizza(pizzaType string) iPizza
	createPizza(pizzaType string) (iPizza, error)
}

// Abstract Concrete Type
type aPizzaStore struct {
	createPizza func(pizzaType string) (iPizza, error)
}

func (a *aPizzaStore) orderPizza(pizzaType string) iPizza {
	if pizza, err := a.createPizza(pizzaType); err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		pizza.prepare()
		pizza.bake()
		pizza.cut()
		pizza.box()
		return pizza
	}
}

/**
 * Each subclass implements the createPizza() method
 * and make use of the orderPizza() method in the parent class
 */
type nyPizzaStore struct {
	*aPizzaStore
}

func newNYPizzaStore() iPizzaStore {
	basePizzaStore := &aPizzaStore{}

	nyPizzaStore := &nyPizzaStore{basePizzaStore}

	nyPizzaStore.aPizzaStore.createPizza = nyPizzaStore.createPizza

	return nyPizzaStore
}

func (n *nyPizzaStore) createPizza(pizzaType string) (iPizza, error) {
	switch pizzaType {
	case "cheese":
		return newNYStyleCheesePizza(), nil
		//case "greek":
		//	return newNYStyleGreekPizza(), nil
		//case "pepperoni":
		//	return newNYStylePepperoniPizza(), nil
	}

	return nil, fmt.Errorf("Invalid pizza type")
}

type chicagoPizzaStore struct {
	*aPizzaStore
}

func newChicagoPizzaStore() iPizzaStore {
	basePizzaStore := &aPizzaStore{}

	chicagoPizzaStore := &chicagoPizzaStore{basePizzaStore}

	chicagoPizzaStore.aPizzaStore.createPizza = chicagoPizzaStore.createPizza

	return chicagoPizzaStore
}

func (c *chicagoPizzaStore) createPizza(pizzaType string) (iPizza, error) {
	switch pizzaType {
	case "cheese":
		return newChicagoStyleCheesePizza(), nil
		//case "greek":
		//	return newChicagoStyleGreekPizza(), nil
		//case "pepperoni":
		//	return newChicagoStylePepperoniPizza(), nil
	}

	return nil, fmt.Errorf("Invalid pizza type")
}
