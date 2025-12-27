package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		i := 1
		for {
			intCh <- i
			i++

			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		i := 1
		for {
			strCh <- "hi" + strconv.Itoa(i)
			i++

			time.Sleep(5 * time.Second)
		}
	}()

	for {
		select {
		case number := <-intCh:
			fmt.Println("Select number", number)
		case str := <-strCh:
			fmt.Println("Select string", str)
		}
	}
}
