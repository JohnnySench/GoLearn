package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var money atomic.Int64
var bank atomic.Int64
var mtx sync.Mutex

func payHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Произошла ошибка чтения body", err)
		return
	}

	requestBodyInt, err := strconv.Atoi(string(requestBody))

	if err != nil {
		fmt.Println("Произошла ошибка конвертации", err)
		return
	}

	mtx.Lock()
	if money.Load()-int64(requestBodyInt) >= 0 {
		time.Sleep(3 * time.Second)
		money.Add(-int64(requestBodyInt))
		fmt.Println("Оплата прошла успешна! Осталось денег:", money.Load())
	} else {
		fmt.Println("Не хватило денег!")
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Произошла ошибка чтения body", err)
		return
	}

	requestBodyInt, err := strconv.Atoi(string(requestBody))

	if err != nil {
		fmt.Println("Произошла ошибка конвертации", err)
		return
	}

	mtx.Lock()
	if money.Load()-int64(requestBodyInt) >= 0 {
		time.Sleep(3 * time.Second)
		money.Add(-int64(requestBodyInt))
		bank.Add(int64(requestBodyInt))
		fmt.Println("Стало денег в банке:", bank.Load())
		fmt.Println("Осталось всего денег:", money.Load())
	} else {
		fmt.Println("Не хватило денег!")
	}
	mtx.Unlock()
}

func main() {
	money.Add(50)

	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	http.ListenAndServe(":9091", nil)
}
