package main

import (
	"fmt"
	"sync"
	"time"
)

var likes int = 0
var rwmtx sync.RWMutex

func setLike(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100_000; i++ {
		rwmtx.Lock()
		likes++
		rwmtx.Unlock()
	}

}

func getLike(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100_000; i++ {
		rwmtx.RLock()
		_ = likes
		rwmtx.RUnlock()
	}
}

func main() {

	wg := sync.WaitGroup{}

	initTime := time.Now()
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go setLike(&wg)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go getLike(&wg)
	}

	wg.Wait()

	fmt.Println("Время выполнения: ", time.Since(initTime), likes)
}
