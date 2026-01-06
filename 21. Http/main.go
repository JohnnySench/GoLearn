package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello, world!"
	bStr := []byte(str)
	_, err := w.Write(bStr)
	if err != nil {
		fmt.Println("Произошла ошибка записи", err.Error())
	} else {
		fmt.Println("Корректно записал")
	}
}

func main() {
	http.HandleFunc("/default", handler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Произошла ошибка", err.Error())
	}

	fmt.Println("Программа закончила свое выполнение")
}
