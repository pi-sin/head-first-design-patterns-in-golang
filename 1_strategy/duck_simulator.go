package main

func main() {
	mallard := NewMallardDuck()
	mallard.PerformFly()
	mallard.PerformQuack()

	rubber := NewRubberDuck()
	rubber.PerformFly()
	rubber.PerformQuack()
}
