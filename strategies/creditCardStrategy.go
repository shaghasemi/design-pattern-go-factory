package strategies

import (
	"bufio"
	"fmt"
	"os"
)

type PayByCreditCard struct {
	reader *bufio.Reader
	card   *CreditCard
}

func NewPayByCreditCard() *PayByCreditCard {
	return &PayByCreditCard{
		reader: bufio.NewReader(os.Stdin),
	}
}
func (p *PayByCreditCard) CollectPaymentDetails() {
	fmt.Println("Enter the card number: ")
	number, _ := p.reader.ReadString('\n')
	number = trimLine(number)
	fmt.Println("Enter the card expiration data 'mm/yy': ")
	date, _ := p.reader.ReadString('\n')
	date = trimLine(date)
	fmt.Println("Enter the cvv code: ")
	cvv, _ := p.reader.ReadString('\n')
	cvv = trimLine(cvv)
	p.card = NewCreditCard(number, date, cvv)
}
func (p *PayByCreditCard) Pay(paymentAmount int) bool {
	if p.cardIsPresent() {
		fmt.Printf("Paying %d using Credit Card.\n", paymentAmount)
		return true
	}
	return false
}
func (p *PayByCreditCard) cardIsPresent() bool {
	return p.card != nil
}
