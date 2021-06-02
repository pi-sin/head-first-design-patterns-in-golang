package main

import "fmt"

func main() {
	// Creating a duck and a turkey
	mallardDuck := &mallardDuck{}

	wildTurkey := &wildTurkey{}

	// Wrapping the turkey in a turkey adapter, which makes it look like a duck
	turkeyAdapter := newTurkeyAdaptor(wildTurkey)

	// Passing the duck to a method testDuck() which expects a Duck object
	fmt.Println("The Duck says...")
	testDuck(mallardDuck)

	// Passing off Turkey as a Duck to the same method
	fmt.Println("\nThe Turkey Adapter says...")
	testDuck(turkeyAdapter)
}

func testDuck(d duck) {
	d.quack()
	d.fly()
}
