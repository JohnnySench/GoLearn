package main

import (
	"study/payments"
	"study/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	cryptoMethod := methods.NewCrypto()

	paymentModule := payments.NewPaymentModule(cryptoMethod)
	idBurger := paymentModule.Pay("Burger", 5)
	idPhone := paymentModule.Pay("Phone", 100)
	paymentModule.Pay("Game", 20)

	info := paymentModule.AllInfo()
	pp.Println(info)

	paymentModule.Cancel(idBurger)

	infoLast := paymentModule.AllInfo()
	pp.Println(infoLast)

	spent := paymentModule.SpendUsd()

	pp.Println(spent)

	infoPhone, _ := paymentModule.Info(idPhone)
	pp.Println(infoPhone)
}
