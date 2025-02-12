package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLine() ([]string, error) {
	fmt.Println("Enter your prices. Confirm every price with Enter")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")

		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)

	return nil
}

func New() CMDManager{
	return CMDManager{}
}
