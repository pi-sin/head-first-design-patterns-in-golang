package main

import "fmt"

func main() {
	// Order up an espresso, no condiments
	beverage1 := &Espresso{}
	fmt.Println("Order1:", beverage1.GetDescription())
	fmt.Println("Cost:", beverage1.Cost())

	// Create a DarkRoast coffee with milk and mocha
	beverage2 := NewMocha(NewMilk(&DarkRoast{}))

	fmt.Println("Order2:", beverage2.GetDescription())
	fmt.Println("Cost:", beverage2.Cost())

	// Add a new decorator (Soy) at runtime
	beverage2 = NewSoy(beverage2)

	fmt.Println("Updated Order2:", beverage2.GetDescription())
	fmt.Println("Updated Cost:", beverage2.Cost())
}
