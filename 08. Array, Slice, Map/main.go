package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	fmt.Println("---------")
	fmt.Println("ARRAYS:")
	Arrays()
	fmt.Println("---------")
	fmt.Println("SLICES:")
	Slices()
	fmt.Println("---------")
	fmt.Println("MAPS:")
	Maps()
}

func Arrays() {
	arr := [...]int{2, 3, 1}
	usersArr := [3]User{
		User{
			Name:    "John",
			Rating:  6.6,
			Premium: true,
		},
		User{
			Name:    "Vika",
			Rating:  6.6,
			Premium: true,
		},
		User{
			Name:    "Alex",
			Rating:  5.0,
			Premium: false,
		},
	}

	arr[0], arr[1] = arr[1], arr[0]

	for i := range arr {
		arr[i] *= 2
	}

	pp.Println(arr)

	for i := range len(arr) {
		arr[i] /= 2
	}

	for i, user := range usersArr {
		if user.Premium {
			usersArr[i].Rating += 1
		}
	}
	pp.Println(usersArr)
}

func Slices() {
	slice := []string{"Hello", "All"}
	slice = append(slice, "World")
	pp.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	intSlice := make([]int, 0, 5)
	intSlice = append(intSlice, 10, 30, 40)
	pp.Println(intSlice)
}

func Maps() {
	weather := map[string]int{
		"03.2025": 3,
		"04.2025": -1,
	}
	weather["05.2025"] = 0
	pp.Println(weather["03.2025"])
	pp.Println(weather)

	for i := range weather {
		weather[i] += 1
	}

	empty, ok := weather["1"]
	fmt.Println(empty, ok)
	pp.Println(weather)

	makeWeather := make(map[string]int, 10)
	fmt.Println(makeWeather)
	fmt.Println(len(makeWeather))
}
