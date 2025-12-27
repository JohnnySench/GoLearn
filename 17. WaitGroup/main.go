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
		fmt.Println("Я почтальон, я отнес газету", text, "в", i, "раз")
	}
}

func main() {
	wq := &sync.WaitGroup{}

	wq.Add(1)
	go postman("Хроники чего-то", wq)

	wq.Add(1)
	go postman("Газета не правда", wq)

	wq.Wait()

	fmt.Println("Main завершился")

}
