package main

import (
	"fmt"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

type SavingsAccount struct {
	balance float64
}

type Customer struct {
	ID   int
	Name string
	Account
}

type Option func(*Customer)

func (s *SavingsAccount) Deposit(amount float64) {
	s.balance += amount
}
func (s *SavingsAccount) Withdraw(amount float64) error {
	if amount > s.balance {
		return fmt.Errorf("amount > balance")
	}
	s.balance -= amount
	return nil
}
func (s *SavingsAccount) Balance() float64 {
	return s.balance
}

func WithName(value string) Option {
	return func(customer *Customer) {
		customer.Name = value
	}
}

func WithAccount(value *SavingsAccount) Option {
	return func(customer *Customer) {
		customer.Account = value
	}
}

func NewCustomer(id int, options ...Option) *Customer {
	customer := &Customer{ID: id}
	for _, option := range options {
		option(customer)
	}
	return customer
}

func main() {
	savings := &SavingsAccount{}
	savings.Deposit(1000)

	customer := NewCustomer(1, WithName("John Doe"), WithAccount(savings))

	err := customer.Account.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Customer: %v, Balance: %v\n", customer.Name, customer.Account.Balance())
}
