package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Payment struct {
	Description string    `json:"description"`
	USD         int       `json:"usd"`
	FullName    string    `json:"fullName"`
	Address     string    `json:"address"`
	Time        time.Time `json:"time"`
}

type HttpResponse struct {
	Money    int       `json:"money"`
	Payments []Payment `json:"payments"`
}

var mtx sync.Mutex
var money = 1000
var history = make([]Payment, 0)

func payHandle(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	var httpResponse HttpResponse
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")
	fmt.Println(fooParam, booParam, "- params query")

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("Err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payment.Time = time.Now()

	mtx.Lock()
	if money-payment.USD > 0 {
		money -= payment.USD
		history = append(history, payment)
	}

	httpResponse = HttpResponse{
		Money:    money,
		Payments: history,
	}

	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("Err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("Err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandle)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Ошибка запуска сервера", err)
	}
}
