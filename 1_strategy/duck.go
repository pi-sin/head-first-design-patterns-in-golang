package main

import "fmt"

type iDuck interface {
	display()
	performFly()
	performQuack()
	setFlyingBehaviour(flyBehaviour)
	setQuackingBehaviour(quackBehaviour)
}

type abstractDuck struct {
	iDuck
	flyingBehaviour   flyBehaviour
	quackingBehaviour quackBehaviour
}

func (d *abstractDuck) display() {
	fmt.Println("Abstract Display")
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

type mallardDuck struct {
	*abstractDuck
}

func newMallardDuck() *mallardDuck {
	d := &abstractDuck{
		flyingBehaviour:   &flyWithWings{},
		quackingBehaviour: &quack{},
	}

	return &mallardDuck{d}
}

func (md *mallardDuck) display() {
	fmt.Println("I’m a real Mallard duck")
}

/**
	Deliberately didn't implement the SetFlyingBehaviour and SetQuackingBehaviour methods
	This is to show how golang works differently than Java
	If we try to use the above methods wrt MallardDuck struct, we will get a run time error unlike a compile time error in JAVA
	The error we get for following line of code is as follows:

	mallardDuck.SetFlyingBehaviour(&FlyRocketPowered{})

	panic: runtime error: invalid memory address or nil pointer dereference
	[signal SIGSEGV: segmentation violation code=0x1 addr=0x30 pc=0x109d60e]

	goroutine 1 [running]:
	main.main()
    	    duck_simulator.go:9 +0xce
*/

type modelDuck struct {
	*abstractDuck
}

func newModelDuck() *modelDuck {
	d := &abstractDuck{
		flyingBehaviour:   &flyNoWay{},
		quackingBehaviour: &muteQuack{},
	}

	return &modelDuck{d}
}

func (md *modelDuck) display() {
	fmt.Println("I’m a real Model duck")
}

func (md *modelDuck) setFlyingBehaviour(fb flyBehaviour) {
	md.flyingBehaviour = fb
}

func (md *modelDuck) setQuackingBehaviour(qb quackBehaviour) {
	md.quackingBehaviour = qb
}
