package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Balance int
}

func Pay(user *User, usd int) error {
	if user.Balance-usd < 0 {
		return errors.New("Error")
	}
	user.Balance -= usd
	return nil
}

type Auto struct {
	Armor int
}

func Gas(a *Auto) (int, error) {
	speed := rand.Intn(150)
	if a.Armor-10 < 0 {
		return speed, errors.New("aвтомобиль может сломаться")
	}
	a.Armor -= 10
	return speed, nil
}

func main() {
	defer func() {
		p := recover() // Возвращает была ли паника
		if p != nil {
			fmt.Println("Произошла паника")
		}
	}()
	slice := []int{1, 2, 3}
	fmt.Println(slice[5])
	user := User{
		Name:    "Johnny",
		Balance: 100,
	}
	pp.Println(user)
	err := Pay(&user, 200)
	if err != nil {
		pp.Println("Произошла ошибка", err.Error())
	}
	pp.Println(user)

	auto := Auto{
		Armor: 25,
	}
	for {
		speed, err := Gas(&auto)
		if err != nil {
			pp.Println("Error:", err.Error())
			break
		}
		pp.Println("Speed:", speed)
	}
}
