package main

import "fmt"

type (
	Human struct{}

	Devil struct {
		Human
		Power string
	}

	Angel struct {
		Human
		WingSpan int
	}
)

func (h Human) Speak() {
	println("Speaking")
}

func (h Human) Walk() {
	println("Walking")
}

func (h Human) Eat() {
	println("Eating")
}

func (d Devil) transform() {
	println("Transforming into Devil with power:", d.Power)
}

func (a Angel) fly() {
	println("Flying with wingspan:", a.WingSpan)
}

func main() {

	human := Human{}
	devil := Devil{
		Power: "Fire",
	}
	angel := Angel{
		WingSpan: 10,
	}

	fmt.Println("Human:")
	human.Speak()
	human.Walk()
	human.Eat()

	fmt.Println("Devil:")
	devil.Speak()
	devil.Walk()
	devil.Eat()
	devil.transform()

	fmt.Println("Angel:")
	angel.Speak()
	angel.Walk()
	angel.Eat()
	angel.fly()

}
