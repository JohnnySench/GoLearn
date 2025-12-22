package payments

type PaymentMethod interface {
	Pay(usd float64) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentsInfo  map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

func (p *PaymentModule) Pay(description string, usd float64) int {
	id := p.paymentMethod.Pay(usd)
	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}
	p.paymentsInfo[id] = info
	return id
}
func (p *PaymentModule) Cancel(id int) {
	p.paymentMethod.Cancel(id)
	if info, ok := p.paymentsInfo[id]; ok {
		info.Cancelled = true
		p.paymentsInfo[id] = info
	}
}
func (p *PaymentModule) Info(id int) (PaymentInfo, int) {
	if info, ok := p.paymentsInfo[id]; ok {
		return info, id
	}
	return PaymentInfo{}, id
}
func (p PaymentModule) AllInfo() map[int]PaymentInfo {
	return p.paymentsInfo
}
func (p PaymentModule) SpendUsd() float64 {
	spend := 0.0
	for _, v := range p.paymentsInfo {
		spend += v.Usd
	}
	return spend
}
