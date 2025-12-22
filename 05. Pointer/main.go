package main

import "fmt"

func main() {
	c := []int{3, 2, 1}
	pointer := &c
	ch(pointer)

	fmt.Println(c)
	fmt.Println(*pointer)

	var ptr *string // nil
	fmt.Println(ptr)

}

func ch(t *[]int) {
	*t = []int{1, 2, 3}
}

// Напишите функцию swapValues, которая принимает два указателя на int
// и меняет их значения местами
func swapValues(a, b *int) {
	a = b
	b = a
}

// Напишите функцию increment, которая принимает указатель на int
// и увеличивает его значение на 1
func increment(n *int) {
	*n++
}

type Person struct {
	Name string
	Age  int
}

// Напишите функцию birthday, которая принимает указатель на Person
// и увеличивает возраст на 1
func birthday(p *Person) {
	p.Age++
}

// Напишите функцию rename, которая принимает указатель на Person
// и новое имя, затем меняет имя персоны
func rename(p *Person, newName string) {
	// Ваш код здесь
}
