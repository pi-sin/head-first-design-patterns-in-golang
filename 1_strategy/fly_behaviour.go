package main

import "fmt"

type flyBehaviour interface {
	fly()
}

type flyWithWings struct{}

func (fw *flyWithWings) fly() {
	fmt.Println("I'm flying")
}

type flyNoWay struct{}

func (fnw *flyNoWay) fly() {
	fmt.Println("I can't fly")
}

type flyRocketPowered struct{}

func (frp *flyRocketPowered) fly() {
	fmt.Println("I’m flying with a rocket!")
}
