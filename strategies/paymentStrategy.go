package strategies

type PaymentStrategy interface {
	Pay(paymentAmount int) bool
	CollectPaymentDetails()
}
