package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Atomic
var number atomic.Int64

func primitiveAtomic() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	go increasePrimitive(&wg)
	go increasePrimitive(&wg)
	go increasePrimitive(&wg)
	go increasePrimitive(&wg)
	go increasePrimitive(&wg)

	go increasePrimitive(&wg)
	go increasePrimitive(&wg)
	go increasePrimitive(&wg)
	go increasePrimitive(&wg)
	go increasePrimitive(&wg)

	wg.Wait()

	fmt.Println("Number:", number.Load())
}

// Mutex
var slice []int
var mtx sync.Mutex

func increasePrimitive(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		number.Add(1)
	}
}
func mutex() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	go increase(&wg)
	go increase(&wg)
	go increase(&wg)
	go increase(&wg)
	go increase(&wg)

	go increase(&wg)
	go increase(&wg)
	go increase(&wg)
	go increase(&wg)
	go increase(&wg)

	wg.Wait()

	fmt.Println("Slice len:", len(slice))
}

func increase(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
	}
}

func main() {
	// primitiveAtomic()
	mutex()
}
