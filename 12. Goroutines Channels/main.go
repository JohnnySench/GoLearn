package main

import (
	"fmt"
	"time"
)

/* не буферезированный канал **/
func mine1(transferPoint chan int, n int) {
	fmt.Println("Поход в шахту номер", n, "начался")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")
	transferPoint <- 10
	fmt.Println("Передал уголь из шахты: ", n)
}

func notBufferChannel() {
	coal := 0

	transferPoint := make(chan int)

	initTime := time.Now()

	go mine1(transferPoint, 1)
	go mine1(transferPoint, 2)
	go mine1(transferPoint, 3)

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	fmt.Println("Добыли угля:", coal)
	fmt.Println("Прошло времени:", time.Since(initTime))
}

/* буферезированный канал **/
func mine2(transferPoint chan int, n int) {
	fmt.Println("Пошел в шахту: ", n)
	time.Sleep(1 * time.Second)
	fmt.Println("Закончил поход в шахту: ", n)
	transferPoint <- 10
	fmt.Println("Передал уголь из шахты: ", n)
}
func bufferChannel() {
	coal := 0

	transferPoint := make(chan int, 3)

	initTime := time.Now()

	go mine2(transferPoint, 1)
	go mine2(transferPoint, 2)
	go mine2(transferPoint, 3)

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	fmt.Println("Угля: ", coal)
	fmt.Println("Прошло времени:", time.Since(initTime))

}

func main() {
	// notBufferChannel()
	bufferChannel()
}
