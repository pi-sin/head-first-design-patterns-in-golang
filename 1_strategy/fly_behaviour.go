package main

import "fmt"

// FlyBehaviour interface that all flying behaviour classes implement
type FlyBehaviour interface {
	Fly()
}

// Flying behaviour implementation for ducks that fly
type FlyWithWings struct{}

// FlyWithWings method
func (fww FlyWithWings) Fly() {
	fmt.Println("I'm flying!")
}

// Flying behaviour implementation for ducks that do not fly
type FlyNoWay struct{}

// FlyNoWay method
func (fnw FlyNoWay) Fly() {
	fmt.Println("I can't fly!")
}

// Rocket-powered flying behaviour
type FlyRocketPowered struct{}

// FlyRocketPowered method
func (frp FlyRocketPowered) Fly() {
	fmt.Println("I'm flying with a rocket!")
}
