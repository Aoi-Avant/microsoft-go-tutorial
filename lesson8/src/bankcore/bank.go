package bank

import (
	"errors"
	"fmt"
)

// interface
type Bank interface {
	Statement()
}

// Customer
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account
type Account struct {
	Customer
	Number  int32
	Balance float64
}

// 預金
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

// 引き出し
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}

	a.Balance -= amount
	return nil
}

// 明細
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

// 送金
func (a *Account) Remittance(toAccount *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if toAccount.Balance < amount {
		return errors.New("the amount to remittance should be greater than the account's balance")
	}

	a.Deposit(amount)
	toAccount.Withdraw(amount)

	return nil
}
