package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	mails := []string{}
	var mtx sync.Mutex
	var coal atomic.Int64
	postmanCtx, postmanCtxCancel := context.WithCancel(context.Background())
	minerCtx, minerCtxCancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}

	go func() {
		time.Sleep(3 * time.Second)
		postmanCtxCancel()
		fmt.Println("--->>> Почтальоны закончили работать!")
	}()

	go func() {
		time.Sleep(6 * time.Second)
		minerCtxCancel()
		fmt.Println("--->>> Шахтеры закончили работать!")
	}()

	postmanCh := postman.PostmanPool(postmanCtx, 3)
	minerCh := miner.MinerPool(minerCtx, 3, true)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range postmanCh {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range minerCh {
			coal.Add(int64(v))
		}
	}()

	wg.Wait()

	fmt.Println("Письма:", len(mails))
	fmt.Println("Уголь:", coal.Load())
}
