package main

import "fmt"

type quackBehaviour interface {
	quack()
}

type quack struct{}

func (q *quack) quack() {
	fmt.Println("Quack")
}

type muteQuack struct{}

func (mq *muteQuack) quack() {
	fmt.Println("Silence")
}

type squeak struct{}

func (s *squeak) quack() {
	fmt.Println("Squeak")
}
