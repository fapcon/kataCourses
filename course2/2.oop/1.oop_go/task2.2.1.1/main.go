package main

import (
	"fmt"
)

type Accounter interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type CurrentAccount struct {
	bal float64
}

type SavingsAccount struct {
	bal float64
}

var err error

func (c *CurrentAccount) Deposit(a float64) error {
	if a > 0.0 {
		c.bal += a
	} else {
		return fmt.Errorf("negative deposit", err)
	}
	return err
}
func (c *CurrentAccount) Withdraw(a float64) error {
	if a <= c.bal {
		c.bal -= a
	} else {
		return fmt.Errorf("withdraw > balance", err)
	}
	return err
}
func (c *CurrentAccount) Balance() float64 {
	return c.bal
}

func (c *SavingsAccount) Deposit(a float64) error {
	if a > 0.0 {
		c.bal += a
	} else {
		return fmt.Errorf("negative deposit", err)
	}
	return err
}
func (c *SavingsAccount) Withdraw(a float64) error {
	if a <= c.bal {
		if c.bal < 500.0 {
			return fmt.Errorf("balance < 500", err)
		} else {
			c.bal -= a
		}
	} else {
		return fmt.Errorf("withdraw > balance", err)
	}
	return err
}
func (c *SavingsAccount) Balance() float64 {
	return c.bal
}

func ProcessAccount(a Accounter) {
	a.Deposit(500)
	a.Withdraw(200)
	fmt.Printf("Balance: %.2f\n", a.Balance())
}

func main() {
	c := &CurrentAccount{bal: 500.00}
	s := &SavingsAccount{bal: 1000.00}
	ProcessAccount(c)
	ProcessAccount(s)
}
