package main

import (
	"fmt"
	"time"
)

type Message struct {
	Author string
	Text   string
}

func main() {
	messageCh1 := make(chan Message)
	messageCh2 := make(chan Message)

	go func() {
		for {
			messageCh1 <- Message{
				Author: "Друг 1",
				Text:   "Привет!",
			}

			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			messageCh2 <- Message{
				Author: "Друг 2",
				Text:   "Как дела?",
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case message1 := <-messageCh1:
			fmt.Println("Сообщение от:", message1.Author, "сообщение:", message1.Text)
		case message2 := <-messageCh2:
			fmt.Println("Сообщение от:", message2.Author, "сообщение:", message2.Text)
		}

	}
}
