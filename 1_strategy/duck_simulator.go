package main

func main() {
	mallardDuck := newMallardDuck()

	mallardDuck.display()
	mallardDuck.performFly()
	mallardDuck.performQuack()
	mallardDuck.swim()

	modelDuck := newModelDuck()

	modelDuck.display()
	modelDuck.performFly()
	modelDuck.performQuack()
	modelDuck.swim()

	modelDuck.setFlyingBehaviour(&flyRocketPowered{})
	modelDuck.performFly()
}
