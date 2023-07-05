package main

func main() {
	mallard := NewMallardDuck()
	mallard.PerformFly()
	mallard.PerformQuack()

	rubber := NewRubberDuck()
	rubber.PerformFly()
	rubber.PerformQuack()

	rubber.SetFlyBehaviour(FlyRocketPowered{})
	rubber.PerformFly()
	rubber.SetQuackBehaviour(MuteQuack{})
	rubber.PerformQuack()
}
