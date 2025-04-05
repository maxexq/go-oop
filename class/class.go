package main

// Defining a struct to represent a class in Go
type Person struct {
	Name string
	Age  int
	Job  string
}

// Defining methods for the Person struct
func (p Person) Greet() {
	println("Hello, my name is", p.Name)
}

func (p Person) GetAge() {
	println("My age is", p.Age)
}

func (p Person) GetJob() {
	println("My job is", p.Job)
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
