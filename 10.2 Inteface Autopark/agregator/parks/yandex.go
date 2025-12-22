package parks

import (
	"fmt"
	"math/rand"
)

type Yandex struct{}

func NewYandex() Yandex {
	return Yandex{}
}

func (p Yandex) AddAuto() int {
	fmt.Println("Яндекс парк добавил новое авто!")
	return rand.Int()
}

func (p Yandex) FixAuto() int {
	fmt.Println("Машина в ремонте")
	return rand.Int()
}

func (p Yandex) DeleteAuto() int {
	fmt.Println("Машина удалена из автопарка Yandex")
	return rand.Int()
}

func (p Yandex) SellAuto() int {
	fmt.Println("Машина продана из автопарка Yandex")
	return rand.Int()
}
