package main

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) error
}

type CreditCard struct {
	balance float64
}

type Bitcoin struct {
	balance float64
}

func (a *CreditCard) Pay(amount float64) error {
	if a.balance > amount {
		a.balance -= amount
		fmt.Printf("Оплачено %v с помощью кредитной карты\n", amount)
	} else {
		return fmt.Errorf("balance < amount")
	}
	return nil
}
func (a *Bitcoin) Pay(amount float64) error {
	if a.balance > amount {
		a.balance -= amount
		fmt.Printf("Оплачено %v с помощью биткоина\n", amount)
	} else {
		return fmt.Errorf("balance < amount")
	}
	return nil
}

func ProcessPayment(p PaymentMethod, amount float64) {
	err := p.Pay(amount)
	if err != nil {
		fmt.Println("Не удалось обработать платеж:", err)
	}
}

func main() {
	cc := &CreditCard{balance: 500.00}
	btc := &Bitcoin{balance: 2.00}

	ProcessPayment(cc, 200.00)
	ProcessPayment(btc, 1.00)
}
