package main

import "fmt"

type turkey interface {
	gobble()
	fly()
}

type wildTurkey struct{}

func (w *wildTurkey) gobble() {
	fmt.Println("Gobble Gobble")
}

func (w *wildTurkey) fly() {
	fmt.Println("I'm flying a short distance")
}
