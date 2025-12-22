package main

import "fmt"

type Auto interface {
	StepOnGas()
}

type BMW struct {
}

type Audi struct {
}

func (a Audi) StepOnGas() {
	fmt.Println("Audi поехало!")
}

func (a BMW) StepOnGas() {
	fmt.Println("BMW поехало!")
}

func ride(a Auto) {
	a.StepOnGas()
}

type Cat struct {
	Name string
}
type Dog struct {
	Name string
}

type Animal interface {
	Run()
}

func (a Dog) Run() {
	fmt.Println("Собака побежала")
}
func (a Cat) Run() {
	fmt.Println("Кошка побежала")
}

func AnimalRun(a Animal) {
	a.Run()
}
func main() {
	bmw := BMW{}
	ride(bmw)

	audi := Audi{}
	ride(audi)

	cat := Cat{Name: "Kesha"}
	dog := Dog{Name: "Sharik"}
	AnimalRun(cat)
	AnimalRun(dog)
}
