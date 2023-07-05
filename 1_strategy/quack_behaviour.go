package main

import "fmt"

// QuackBehavior interface that all quacking behaviour classes implement
type QuackBehaviour interface {
	Quack()
}

// QuackBehavior implementation for ducks that quack
type Quack struct{}

// Quack method
func (q Quack) Quack() {
	fmt.Println("Quack!")
}

// QuackBehavior implementation for ducks that squeak
type Squeak struct{}

// Squeak method
func (s Squeak) Quack() {
	fmt.Println("Squeak!")
}

// QuackBehavior implementation for ducks that don't quack
type MuteQuack struct{}

// MuteQuack method
func (mq MuteQuack) Quack() {
	fmt.Println("Silence!")
}
