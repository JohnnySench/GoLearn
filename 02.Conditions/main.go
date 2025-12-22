package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	score := -32

	if score >= 1 {
		fmt.Println("score 1 || score > 1")
	}

	fmt.Println(strconv.Itoa(int(math.Abs(float64(score)))))

	fmt.Println(-3 % 10)
}
