package payment

import "errors"

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	payments      map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(method PaymentMethod) *PaymentModule {
	return &PaymentModule{
		payments:      make(map[int]PaymentInfo),
		paymentMethod: method,
	}
}

func (p PaymentModule) Pay(descr string, usd int) int {
	id := p.paymentMethod.Pay(usd)
	info := PaymentInfo{
		Description: descr,
		Usd:         usd,
		Cancelled:   false,
	}
	p.payments[id] = info
	return id
}

func (p PaymentModule) Cancel(id int) (PaymentInfo, error) {
	if payment, ok := p.payments[id]; ok {
		payment.Cancelled = true
		p.payments[id] = payment
		return payment, nil
	}
	return PaymentInfo{}, errors.New("не нашли такой операции")
}

func (p PaymentModule) Info(id int) (PaymentInfo, error) {
	if payment, ok := p.payments[id]; ok {
		return payment, nil
	}
	return PaymentInfo{}, errors.New("нет такого")
}

func (p PaymentModule) AllInfo() map[int]PaymentInfo {
	return p.payments
}
