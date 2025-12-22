package methods

import (
	"fmt"
	"math/rand"
)

type Paypal struct {
}

func NewPaypal() Paypal {
	return Paypal{}
}

func (p Paypal) Pay(usd float64) int {
	fmt.Println("Оплата Paypal")
	return rand.Int()
}

func (p Paypal) Cancel(id int) {
	fmt.Println("Отмена Paypal-операции", id)
}
