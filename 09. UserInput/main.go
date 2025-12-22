package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	HandleInput()
}

func HandleInput() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите команду")

	if ok := scanner.Scan(); !ok {
		fmt.Println("Ошибка ввода")
		return
	}

	text := scanner.Text()
	fields := strings.Fields(text)
	cmd := fields[0]

	if cmd == "добавить" {
		fmt.Println("Вы хотите добавить:", strings.Join(fields[1:], ", "))
	}
	fmt.Println("Text:", text)
	fmt.Println(strings.Fields(text))
}
