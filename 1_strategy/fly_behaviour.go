package main

import "fmt"

/**
 * Interface that all flying behaviour classes implement
 */
type flyBehaviour interface {
	fly()
}

/**
 * Flying behaviour implementation for ducks that fly
 */
type flyWithWings struct{}

func (fw *flyWithWings) fly() {
	fmt.Println("I'm flying")
}

/**
 * Flying behaviour implementation for ducks
 * that do not fly (like rubber and decoy ducks)
 */
type flyNoWay struct{}

func (fnw *flyNoWay) fly() {
	fmt.Println("I can't fly")
}

/**
 * Rocket-powered flying behaviour
 */
type flyRocketPowered struct{}

func (frp *flyRocketPowered) fly() {
	fmt.Println("I’m flying with a rocket!")
}
