package main

import (
	"study/agregator"
	"study/agregator/parks"

	"github.com/k0kubun/pp"
)

func main() {
	yandex := parks.NewYandex()

	park := agregator.NewAgregator(yandex)

	id := park.AddAuto("Bmw X5", "BMW", 2000000)
	autos := park.AllAutos()
	pp.Println(autos)
	park.DeleteAuto(id)
	autos1 := park.AllAutos()
	pp.Println(autos1)
}
