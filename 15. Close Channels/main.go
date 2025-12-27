package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
Канал закрываем один раз и больше не работаем с ним
При чтении закрытого канал возвращается default value value, ok := ch
*/

func base() {
	// Создаем открытый канал
	ch := make(chan int, 1)

	ch <- 1
	// Закрываем канал
	close(ch)

	// Читаю значение из закрытого канала
	v1, ok1 := <-ch
	v2, ok2 := <-ch
	v3, ok3 := <-ch

	fmt.Println("v:", v1, v2, v3)
	fmt.Println("ok:", ok1, ok2, ok3)

	// Запись в закрытый канал (Panica)
	ch <- 10
}

func closeChannel() {

	transferPoint := make(chan int)
	coal := 0

	// fullVersion := func() {
	// 	for {
	// 		v, ok := <-transferPoint
	// 		if !ok {
	// 			fmt.Println("Шахтер добыл весь уголь")
	// 			break
	// 		}
	// 		coal += v
	// 		fmt.Println("Угля:", coal)
	// 	}
	// }

	shortVersion := func() {
		for v := range transferPoint {
			coal += v
			fmt.Println("Угля:", coal)
		}
	}

	go func() {
		iterations := 3 + rand.Intn(4)
		fmt.Println("Пошел в шахту...")
		fmt.Println("Итераций:", iterations)

		for i := 1; i <= iterations; i++ {
			time.Sleep(1 * time.Second)
			transferPoint <- 10
		}
		close(transferPoint)
	}()

	// Длинная запись похода в канал
	// fullVersion()

	// Короткая запись похода в канал
	shortVersion()

	fmt.Println("Всего добыто угля:", coal)
}

func main() {
	// base()
	closeChannel()
}
