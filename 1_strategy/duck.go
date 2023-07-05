package main

// Declare two reference variables for the behaviour interface types.
// All duck subclasses inherit these.
type Duck struct {
	flyBehaviour   FlyBehaviour
	quackBehaviour QuackBehaviour
}

// Delegate fly and quack behaviour to respective behaviour classes
func (d Duck) PerformFly() {
	d.flyBehaviour.Fly()
}

func (d Duck) PerformQuack() {
	d.quackBehaviour.Quack()
}

// Setting/Changing the behaviour dynamically
func (d *Duck) SetFlyBehaviour(fb FlyBehaviour) {
	d.flyBehaviour = fb
}

func (d *Duck) SetQuackBehaviour(qb QuackBehaviour) {
	d.quackBehaviour = qb
}

// MallardDuck struct
type MallardDuck struct {
	Duck
}

// MallardDuck constructor
func NewMallardDuck() *MallardDuck {
	mallard := &MallardDuck{}
	mallard.flyBehaviour = FlyWithWings{}
	mallard.quackBehaviour = Quack{}
	return mallard
}

// RubberDuck struct
type RubberDuck struct {
	Duck
}

// RubberDuck constructor
func NewRubberDuck() *RubberDuck {
	rubber := &RubberDuck{}
	rubber.flyBehaviour = FlyNoWay{}
	rubber.quackBehaviour = Squeak{}
	return rubber
}
