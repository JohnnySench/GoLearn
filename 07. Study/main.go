package main

import (
	"fmt"
	"study/greeting"
	"study/user"

	"github.com/k0kubun/pp"
)

func main() {
	fmt.Println("Main")
	greeting.SayHello()
	user := user.NewUser(
		28,
		"Johnny",
	)
	fmt.Println(user)
	user.SetAge(32)
	user.SetName("John")
	fmt.Println(user)
	pp.Println(user)
}
