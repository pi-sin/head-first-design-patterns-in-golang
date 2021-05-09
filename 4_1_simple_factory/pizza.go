package main

import "fmt"

type iPizza interface {
	prepare()
	bake()
	cut()
	box()
}

type pizza struct {
	pizzaType string
}

func (p *pizza) prepare() {
	fmt.Printf("Preparing %s Pizza\n", p.pizzaType)
}

func (p *pizza) bake() {
	fmt.Printf("Baking %s Pizza\n", p.pizzaType)
}

func (p *pizza) cut() {
	fmt.Printf("Cutting %s Pizza\n", p.pizzaType)
}

func (p *pizza) box() {
	fmt.Printf("Boxing %s Pizza\n", p.pizzaType)
}

type cheesePizza struct {
	*pizza
}

func newCheesePizza() iPizza {
	p := &pizza{
		pizzaType: "Cheese",
	}

	return &cheesePizza{
		pizza: p,
	}
}

type greekPizza struct {
	*pizza
}

func newGreekPizza() iPizza {
	p := &pizza{
		pizzaType: "Greek",
	}

	return &greekPizza{
		pizza: p,
	}
}

type pepperoniPizza struct {
	*pizza
}

func newPepperoniPizza() iPizza {
	p := &pizza{
		pizzaType: "Pepperoni",
	}

	return &pepperoniPizza{
		pizza: p,
	}
}
