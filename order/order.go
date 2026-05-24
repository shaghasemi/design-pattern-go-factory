package order

import "design-pattern-go-factory/strategies"

type Order struct {
	totalCost int
	isClosed  bool
}

func NewOrder() *Order {
	return &Order{}
}
func (o *Order) ProcessOrder(strategy strategies.PaymentStrategy) {
	strategy.CollectPaymentDetails()
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
