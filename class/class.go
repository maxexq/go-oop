package main

import "fmt"

// Defining a struct to represent a class in Go
type Person struct {
	Name string
	Age  int
	Job  string
}

// Defining methods for the Person struct
func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func (p Person) GetAge() {
	fmt.Println("My age is", p.Age)
}

func (p Person) GetJob() {
	fmt.Println("My job is", p.Job)
}

func main() {
	// Creating an instance of the Person struct
	softwareEngineer := Person{
		Name: "John Doe",
		Age:  30,
		Job:  "Software Engineer",
	}

	uxDesigner := Person{
		Name: "Jane Smith",
		Age:  28,
		Job:  "UX Designer",
	}

	softwareEngineer.Greet()
	softwareEngineer.GetAge()
	softwareEngineer.GetJob()

	uxDesigner.Greet()
	uxDesigner.GetAge()
	uxDesigner.GetJob()

}
