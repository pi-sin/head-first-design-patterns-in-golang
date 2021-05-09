package main

func main() {
	nyPizzaFactory := &nyPizzaFactory{}

	pizzaStore := newPizzaStore(nyPizzaFactory)
	pizzaStore.orderPizza("pepperoni")
	pizzaStore.orderPizza("new york cheese")
	pizzaStore.orderPizza("greek")
}
