package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Отнес сообщение:", text, "в", i, "раз")
	}
}

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)
	go postman("как дела?", &wg)

	wg.Add(1)
	go postman("Привет!", &wg)

	wg.Wait()

}
