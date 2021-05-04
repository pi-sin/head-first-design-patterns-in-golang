package main

import "fmt"

func main() {
	/**
	 * Order up an espresso, no condiments
	 * and print its description and cost
	 */
	beverage := &espresso{}

	fmt.Printf("%s $%.2f\n", beverage.description(), beverage.cost())

	// Make a DarkRoast object.
	darkRoast := &darkRoast{}

	// Wrap it with a Mocha
	singleMocha := &mocha{
		beverage: darkRoast,
	}

	// Wrap it in a second Mocha
	doubleMocha := &mocha{
		beverage: singleMocha,
	}

	// Wrap it in a Whip.
	doubleMochaWhip := &whip{
		beverage: doubleMocha,
	}

	fmt.Printf("%s $%.2f\n", doubleMochaWhip.description(), doubleMochaWhip.cost())

	// Finally give us a HouseBlend with Soy, Mocha, and Whip.
	soyMochaWhipHouseBlend := &whip{
		beverage: &mocha{
			beverage: &soy{
				beverage: &houseBlend{},
			},
		},
	}

	fmt.Printf("%s $%.2f\n", soyMochaWhipHouseBlend.description(), soyMochaWhipHouseBlend.cost())
}
