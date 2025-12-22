package main

import (
	"fmt"
	"time"
)

func calc(arr []int) int {
	mapa := make(map[int]int, len(arr))

	for _, v := range arr {
		mapa[v]++
	}
	return mapa[mapa[1]]
}

func main() {
	stant := []int{1, 4, 3, 2, 1, 2, 3, 2}
	res := calc(stant)
	fmt.Println(res)

	score := 0
	for {
		score++
		time.Sleep(500 * time.Millisecond)
		if score == 10 {
			break
		}
	}
}
