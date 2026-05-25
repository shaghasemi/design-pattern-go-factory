package strategies

type PayStrategy interface {
	Pay(paymentAmount int) bool
	CollectPaymentDetails()
}
