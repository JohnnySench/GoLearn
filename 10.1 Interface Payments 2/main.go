package main

import (
	"study/payment"
	"study/payment/methods"

	"github.com/k0kubun/pp"
)

func main() {
	paypal := methods.NewPaypal()
	module := payment.NewPaymentModule(paypal)
	module.Pay("Auto", 200)
	payment, err := module.Cancel(12)

	pp.Println(payment)
	pp.Println(err.Error())

}
