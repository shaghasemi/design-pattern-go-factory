package strategies

type CreditCard struct {
	amount int
	number string
	data   string
	cvv    string
}

func NewCreditCard(number, data, cvv string) *CreditCard {
	return &CreditCard{
		amount: 100000,
		number: number,
		data:   data,
		cvv:    cvv,
	}
}

func (c *CreditCard) SetAmount(amount int) {
	c.amount = amount
}

func (c *CreditCard) Amount() int {
	return c.amount
}
