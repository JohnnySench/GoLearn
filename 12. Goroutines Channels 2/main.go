package main

import (
	"fmt"
	"time"
)

func mine(channel chan int, n int) {
	fmt.Println("Поход в шахту номер: ", n)
	time.Sleep(time.Second * 1)
	fmt.Println("Начинаю передавать уголь: ", n)

	channel <- 10

	fmt.Println("Поход закончен: ", n)

	// ...
}

func noBuffer() {
	coal := 0

	transferPoint := make(chan int)

	initialTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	fmt.Println("Добыли угля: ", coal)
	fmt.Println("Заняло времени: ", time.Since(initialTime))
}

func buffer() {
	coal := 0

	transferPoint := make(chan int, 3)

	initialTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	fmt.Println("Добыли угля: ", coal)
	fmt.Println("Заняло времени: ", time.Since(initialTime))
}

func main() {
	noBuffer()
	// buffer()
}
