package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("Defer 1")
	}()
	defer func() {
		fmt.Println("Defer 2")
	}()

	fmt.Println("Main")
	defer square(4)
}

func square(n int) {
	fmt.Println(n * n)
}
