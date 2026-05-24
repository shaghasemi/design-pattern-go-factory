package main

import (
	"bufio"
	"design-pattern-go-factory/order"
	"design-pattern-go-factory/strategies"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var priceOfProducts = map[int]int{
	1: 2200,
	2: 1850,
	3: 1100,
	4: 890,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	order := order.NewOrder()
	var strategy strategies.PaymentStrategy

	for !order.IsClosed() {
		var continueChoice string

		for {
			fmt.Println("Please, select a product:")
			fmt.Println("1 - Mother board")
			fmt.Println("2 - CPU")
			fmt.Println("3 - HDD")
			fmt.Println("4 - Memory")

			choice, _ := readInt(reader)
			cost := priceOfProducts[choice]

			fmt.Println("Count: ")
			count, _ := readInt(reader)

			order.SetTotalCost(cost * count)

			fmt.Println("Do you wish to continue selecting products? Y/N: ")
			continueChoice, _ = reader.ReadString('\n')
			continueChoice = strings.TrimSpace(continueChoice)

			if !strings.EqualFold(continueChoice, "Y") {
				break
			}
		}

		if strategy == nil {
			fmt.Println("Please, select a payment method:")
			fmt.Println("1 - PalPay")
			fmt.Println("2 - Credit Card")

			paymentMethod, _ := reader.ReadString('\n')
			paymentMethod = strings.TrimSpace(paymentMethod)

			if paymentMethod == "1" {
				strategy = strategies.NewPayPalStrategy()
			} else {
				strategy = strategies.NewPayByCreditCard()
			}
		}

		order.ProcessOrder(strategy)

		fmt.Printf("Pay %d units or Continue shopping? P/C: ", order.TotalCost())
		proceed, _ := reader.ReadString('\n')
		proceed = strings.TrimSpace(proceed)

		if strings.EqualFold(proceed, "P") {
			if strategy.Pay(order.TotalCost()) {
				fmt.Println("Payment has been successful.")
			} else {
				fmt.Println("FAIL! Please, check your data.")
			}
			order.SetClosed()
		}
	}

}

func readInt(reader *bufio.Reader) (int, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.TrimSpace(line))
}
