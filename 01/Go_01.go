package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func (h Human) Greet() {
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", h.Name, h.Age)
}

func (h Human) GetAge() int {
	return h.Age
}

func (h Human) Walk() {
	fmt.Printf("%s is walking.\n", h.Name)
}

func NewHuman(name string, age int) Human {
	return Human{Name: name, Age: age}
}

type Action struct {
	Human
	Profession string
}

func (a Action) PerformAction() {
	fmt.Printf("%s is performing %s.\n", a.Name, a.Profession)
}

func (a Action) Greet() {
	fmt.Printf("Hello, my name is %s and I'm %d years old. I am also a %s.\n", a.Name, a.Age, a.Profession)
}

func main() {
	worker := Action{
		Human: Human{Name: "John Doe",
			Age: 30},
		Profession: "Gopher",
	}
	worker.Human.Greet()
	worker.Greet()
	worker.PerformAction()
	worker.Walk()
	worker.Human = NewHuman("Samsonlu", 31)
	worker.Human.Greet()
	worker.Greet()
	worker.PerformAction()
	worker.Walk()
}
