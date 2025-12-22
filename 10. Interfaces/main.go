package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type Auto interface {
	StepOnGas()
}

type BMW struct {
	Price int
}

type Audi struct {
	Price int
}

type Lada struct {
	Price int
}

func (a Audi) StepOnGas() {
	fmt.Println("Audi поехало!")
}

func (a BMW) StepOnGas() {
	fmt.Println("BMW поехало!")
}

func (a Lada) StepOnGas() {
	fmt.Println("Lada кое как поехала")
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

type Human interface {
	Go()
	Eat()
}
type Jenya struct {
	Age int
}
type Vika struct {
	Age int
}

func (h Jenya) Go() {
	fmt.Println("Я человек и мне:", h.Age)
}
func (h Vika) Go() {
	fmt.Println("Я человек и мне:", h.Age)
}
func (h Jenya) Eat() {
	fmt.Println("Я человек и я ем")
}
func (h Vika) Eat() {
	fmt.Println("Я человек и я ем")
}
func HelloHuman(h Human) {
	h.Go()
	h.Eat()
}
func main() {
	bmw := BMW{Price: 100}
	ride(bmw)

	audi := Audi{Price: 90}
	ride(audi)

	lada := Lada{Price: 10}
	ride(lada)

	autos := []Auto{}

	autos = append(autos, bmw, audi, lada)

	pp.Println(autos)

	cat := Cat{Name: "Kesha"}
	dog := Dog{Name: "Sharik"}
	AnimalRun(cat)
	AnimalRun(dog)

	jenya := Jenya{Age: 28}
	vika := Vika{Age: 22}
	HelloHuman(jenya)
	HelloHuman(vika)
}
