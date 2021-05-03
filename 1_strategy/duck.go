package main

import "fmt"

type abstractDuck struct {
	display func()
	/**
	 * Declare two reference variables for the
	 * behaviour interface types.
	 * All duck subclasses inherit these
	 */
	flyingBehaviour   flyBehaviour
	quackingBehaviour quackBehaviour
}

/**
 * Delegate fly and quack behaviour to
 * respective behaviour classes
 */
func (d *abstractDuck) performFly() {
	d.flyingBehaviour.fly()
}

func (d *abstractDuck) performQuack() {
	d.quackingBehaviour.quack()
}

func (d *abstractDuck) swim() {
	fmt.Println("All ducks float, even decoys!")
}

/**
 * Setting the behaviour dynamically
 */
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
		// implementing the abstract method
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
		// implementing the abstract method
		display: func() {
			fmt.Println("I’m a real Model duck")
		},
		flyingBehaviour:   &flyNoWay{},
		quackingBehaviour: &muteQuack{},
	}

	return &modelDuck{d}
}
