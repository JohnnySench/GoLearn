package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(
	ctx context.Context,
	ch chan<- int,
	wg *sync.WaitGroup,
	n int,
	power int,
	immediate bool,
) {
	defer wg.Done()
	for {
		if immediate {
			fmt.Println("Я шахтер номер номер:", n, "начал добывать угль")
			select {
			case <-ctx.Done():
				fmt.Println("Я шахтер номер:", n, "и я закончил работу!")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Я шахтер номер:", n, "Добыл угля:", power)
			}
			select {
			case <-ctx.Done():
				fmt.Println("Я шахтер номер:", n, "и я закончил работу!")
				return
			case ch <- power:
				fmt.Println("Я шахтер номер:", n, "Передал:", power, "угля")
			}
		} else {
			select {
			case <-ctx.Done():
				fmt.Println("Я шахтер номер:", n, "и я закончил работу!")
				return
			default:
				fmt.Println("Я шахтер номер номер:", n, "начал добывать угль")
				time.Sleep(1 * time.Second)
				fmt.Println("Я шахтер номер:", n, "Добыл угля:", power)

				ch <- power
				fmt.Println("Я шахтер номер:", n, "Передал:", power, "угля")
			}
		}

	}
}

func MinerPool(ctx context.Context, minerCount int, immediate bool) <-chan int {
	chMiner := make(chan int)

	wg := sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go miner(ctx, chMiner, &wg, i, i*10, immediate)
	}

	go func() {
		wg.Wait()
		close(chMiner)
	}()

	return chMiner
}
