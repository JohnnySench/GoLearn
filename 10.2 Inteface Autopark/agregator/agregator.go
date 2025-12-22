package agregator

type Park interface {
	AddAuto() int
	FixAuto() int
	DeleteAuto() int
	SellAuto() int
}

type Agregator struct {
	autos map[int]Auto
	park  Park
}

func NewAgregator(park Park) *Agregator {
	return &Agregator{
		autos: make(map[int]Auto),
		park:  park,
	}
}

func (a *Agregator) AddAuto(name string, company string, price float64) int {
	idAuto := a.park.AddAuto()
	auto := Auto{
		Name:    name,
		Company: company,
		Price:   price,
	}
	a.autos[idAuto] = auto
	return idAuto
}

func (a *Agregator) FixAuto(idAuto int) {
	a.park.FixAuto()
	if auto, ok := a.autos[idAuto]; ok {
		auto.NeedFix = true
		a.autos[idAuto] = auto
	}
}

func (a *Agregator) SellAuto(idAuto int) {
	a.park.SellAuto()
	if auto, ok := a.autos[idAuto]; ok {
		auto.IsSell = true
		a.autos[idAuto] = auto
	}
}

func (a *Agregator) DeleteAuto(idAuto int) {
	a.park.DeleteAuto()
	delete(a.autos, idAuto)
}

func (a Agregator) AllAutos() map[int]Auto {
	return a.autos
}
