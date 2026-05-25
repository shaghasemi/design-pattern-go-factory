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
	o := order.NewOrder()

	for !o.IsClosed() {
		// var continueChoice string

		for {
			fmt.Println("Please, select a product:")
			fmt.Println("1 - Mother board")
			fmt.Println("2 - CPU")
			fmt.Println("3 - HDD")
			fmt.Println("4 - Memory")

			choice, err := readInt(reader)
			if err != nil {
				fmt.Println("Invalid number, try again.")
				continue
			}

			cost, ok := priceOfProducts[choice]

			if !ok {
				fmt.Println("Unknown product, try again.")
				continue
			}

			fmt.Println("Count: ")
			count, err := readInt(reader)
			if err != nil {
				fmt.Println("Invalid count, try again.")
				continue
			}

			o.SetTotalCost(cost * count)

			fmt.Println("Do you wish to continue selecting products? Y/N: ")
			continueChoice, _ := reader.ReadString('\n')
			continueChoice = strings.TrimSpace(continueChoice)

			if !strings.EqualFold(continueChoice, "Y") {
				break
			}
		}

		if !o.HasStrategy() {
			fmt.Println("Please, select a payment method:")
			fmt.Println("1 - PalPay")
			fmt.Println("2 - Credit Card")

			paymentMethod, _ := reader.ReadString('\n')
			paymentMethod = strings.TrimSpace(paymentMethod)

			switch paymentMethod {
			case "1":
				o.SetPayStrategy(strategies.NewPayPalStrategy(reader))
			case "2":
				o.SetPayStrategy(strategies.NewPayPalStrategy(reader))
			default:
				fmt.Println("Invalid choice, try again.")
			}

		}

		o.ProcessOrder()

		fmt.Printf("Pay %d units or Continue shopping? P/C: ", o.TotalCost())
		proceed, _ := reader.ReadString('\n')
		proceed = strings.TrimSpace(proceed)

		if strings.EqualFold(proceed, "P") {
			if o.Pay() {
				fmt.Println("Payment has been successful.")
			} else {
				fmt.Println("FAIL! Please, check your data.")
			}

			o.SetClosed()
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
