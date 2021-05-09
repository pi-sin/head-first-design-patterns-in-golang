package main

import "fmt"

func main() {
	nyPizzaStore := newNYPizzaStore()
	chicagoPizzaStore := newChicagoPizzaStore()

	pizza := nyPizzaStore.orderPizza("cheese")

	fmt.Printf("Ethan ordered %s pizza\n\n", pizza.getName())

	pizza = chicagoPizzaStore.orderPizza("cheese")

	fmt.Printf("Joel ordered %s pizza\n\n", pizza.getName())

}
