package main

import "fmt"

/**
 * We will start with the abstract implement of Pizza.
 * All concrete Pizzas will derive from this.
 */
type iPizza interface {
	prepare()
	bake()
	cut()
	box()
	getName() string
}

type pizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

// Implementing the default preparation steps for pizza: prepare, bake, cut, box
func (p *pizza) prepare() {
	fmt.Printf("Preparing %s Pizza\n", p.name)
	fmt.Printf("Tossing %s dough\n", p.dough)
	fmt.Printf("Adding %s sauce\n", p.sauce)
	fmt.Print("Adding toppings: \n")
	for _, t := range p.toppings {
		fmt.Printf("\t%s, ", t)
	}
	fmt.Println()
}

func (p *pizza) bake() {
	fmt.Println("Baking Pizza for 25 min at 350")
}

func (p *pizza) cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p *pizza) box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (p *pizza) getName() string {
	return p.name
}

type nyStyleCheesePizza struct {
	*pizza
}

func newNYStyleCheesePizza() iPizza {
	p := &pizza{
		name:     "NY Style Sauce and Cheese Pizza",
		dough:    "Thin Crust Dough",
		sauce:    "Marinara Sauce",
		toppings: []string{"Grated Reggiano Cheese"},
	}

	return &nyStyleCheesePizza{
		pizza: p,
	}
}

type chicagoStyleCheesePizza struct {
	*pizza
}

func newChicagoStyleCheesePizza() iPizza {
	p := &pizza{
		name:     "Chicago Style Deep Dish Cheese Pizza",
		dough:    "Extra Thick Crust Dough",
		sauce:    "Plum Tomato Sauce",
		toppings: []string{"Shredded Mozzarella Cheese"},
	}

	return &chicagoStyleCheesePizza{
		pizza: p,
	}
}

/**
 * The Chicago Style Pizza overrides the cut method
 * so that the pieces are cut in squares.
 */
func (c *chicagoStyleCheesePizza) cut() {
	fmt.Println("Cutting the pizza into square slices")
}
