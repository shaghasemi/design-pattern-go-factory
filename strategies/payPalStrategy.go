package strategies

import (
	"bufio"
	"fmt"
	"os"
)

var payPalDatabase = map[string]string{
	"amanda1985": "amanda@ya.com",
	"qwerty":     "john@amazon.eu",
}

type PayPalStrategy struct {
	reader   *bufio.Reader
	email    string
	password string
	signedIn bool
}

func NewPayPalStrategy() *PayPalStrategy {
	return &PayPalStrategy{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (p *PayPalStrategy) CollectPaymentDetails() {
	for !p.signedIn {
		fmt.Println("Enter the user's email: ")
		email, _ := p.reader.ReadString('\n')
		p.email = trimLine(email)

		fmt.Println("Enter the password: ")
		password, _ := p.reader.ReadString('\n')
		p.password = trimLine(password)

		if p.verify() {
			fmt.Println("Data verification has been successful.")
		} else {
			fmt.Println("Wrong email or password")
		}
	}

}

func (p *PayPalStrategy) verify() bool {
	p.signedIn = p.email == payPalDatabase[p.password]
	return p.signedIn
}

func (p *PayPalStrategy) Pay(paymentAmount int) bool {
	if p.signedIn {
		fmt.Printf("Paying %d using PayPal. \n", paymentAmount)
		return true
	} else {
		return false
	}
}
