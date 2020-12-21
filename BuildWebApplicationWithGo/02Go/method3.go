package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Studnet struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

func main() {
	mark := Studnet{Human{"Mark", 25, "222-222-yyyy"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "333-222-xxx"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}
