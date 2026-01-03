package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(
	ctx context.Context,
	ch chan<- string,
	wg *sync.WaitGroup,
	n int,
	mail string,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Рабочий день почтальона номер:", n, "закончился!!!")
			return
		default:
			fmt.Println("Я почтальон номер:", n, "Взял письмо")
			time.Sleep(1 * time.Second)
			fmt.Println("Я почтальон номер:", n, "Отнес письмо до почты	", mail)

			ch <- mail
			fmt.Println("Я почтальон номер:", n, "Передал письмо!", mail)
		}
	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	chPostman := make(chan string)
	wg := sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go postman(ctx, chPostman, &wg, i, getMail(i))
	}

	go func() {
		wg.Wait()
		close(chPostman)
	}()

	return chPostman
}

func getMail(n int) string {
	mails := map[int]string{
		1: "Северный рабочий",
		2: "New York time`s",
		3: "Северодвинск Life",
	}

	v, ok := mails[n]
	if !ok {
		return "Лотерея"
	}
	return v
}
