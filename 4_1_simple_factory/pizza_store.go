package main

import (
	"fmt"
)

/**
 * PizzaStore now has a reference to SimplePizzaFactory
 */
type pizzaStore struct {
	factory pizzaFactory
}

/**
 * The Pizza Store constructor now takes the reference of the pizza factory
 */
func newPizzaStore(factory pizzaFactory) *pizzaStore {
	return &pizzaStore{
		factory: factory,
	}
}

/**
 * We are passing the type of pizza
 */
func (ps *pizzaStore) orderPizza(pizzaType string) {
	/**
	 * Now we pass the type to the factory.
	 * No more concrete instantiations here!
	 */
	if pizza, err := ps.factory.createPizza(pizzaType); err != nil {
		fmt.Println(err.Error())
	} else {
		pizza.prepare()
		pizza.bake()
		pizza.cut()
		pizza.box()
	}
}