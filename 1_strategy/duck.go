package main

import "fmt"

type abstractDuck struct {
	display           func()
	flyingBehaviour   flyBehaviour
	quackingBehaviour quackBehaviour
}

func (d *abstractDuck) performFly() {
	d.flyingBehaviour.fly()
}

func (d *abstractDuck) performQuack() {
	d.quackingBehaviour.quack()
}

func (d *abstractDuck) swim() {
	fmt.Println("All ducks float, even decoys!")
}

func (d *abstractDuck) setFlyingBehaviour(fb flyBehaviour) {
	d.flyingBehaviour = fb
}

func (d *abstractDuck) setQuackingBehaviour(qb quackBehaviour) {
	d.quackingBehaviour = qb
}

type mallardDuck struct {
	*abstractDuck
}

func newMallardDuck() *mallardDuck {
	d := &abstractDuck{
		display: func() {
			fmt.Println("I’m a real Mallard duck")
		},
		flyingBehaviour:   &flyWithWings{},
		quackingBehaviour: &quack{},
	}

	return &mallardDuck{d}
}

type modelDuck struct {
	*abstractDuck
}

func newModelDuck() *modelDuck {
	d := &abstractDuck{
		display: func() {
			fmt.Println("I’m a real Model duck")
		},
		flyingBehaviour:   &flyNoWay{},
		quackingBehaviour: &muteQuack{},
	}

	return &modelDuck{d}
}
