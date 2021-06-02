package main

import "fmt"

type duck interface {
	quack()
	fly()
}

type mallardDuck struct{}

func (m *mallardDuck) quack() {
	fmt.Println("Quack")
}

func (m *mallardDuck) fly() {
	fmt.Println("I'm flying")
}
