package methods

import (
	"fmt"
	"math/rand"
)

type Paypal struct {
}

func NewPaypal() *Paypal {
	return &Paypal{}
}

func (p *Paypal) Pay(usd int) int {
	fmt.Println("Оплатил с помощью Paypal", usd)
	return rand.Int()
}

func (p *Paypal) Cancel(id int) {
	fmt.Println("Отменил операцию", id)
}
