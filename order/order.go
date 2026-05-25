package order

import "design-pattern-go-factory/strategies"

type Order struct {
	strategy                strategies.PayStrategy
	paymentDetailsCollected bool
	totalCost               int
	isClosed                bool
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) ProcessOrder() {
	if o.strategy == nil {
		return
	}
	if o.paymentDetailsCollected {
		return
	}
	o.strategy.CollectPaymentDetails()
	o.paymentDetailsCollected = true
}

func (o *Order) SetTotalCost(cost int) {
	o.totalCost += cost
}

func (o *Order) TotalCost() int {
	return o.totalCost
}

func (o *Order) IsClosed() bool {
	return o.isClosed
}

func (o *Order) SetClosed() {
	o.isClosed = true
}

func (o *Order) SetPayStrategy(strategy strategies.PayStrategy) {
	o.strategy = strategy
	o.paymentDetailsCollected = false
}

func (o *Order) Pay() bool {
	if o.strategy == nil {
		return false
	}
	return o.strategy.Pay(o.totalCost)
}

func (o *Order) HasStrategy() bool {
	return o.strategy != nil
}
