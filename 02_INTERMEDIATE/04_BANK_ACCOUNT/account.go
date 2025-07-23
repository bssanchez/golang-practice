package account

import (
	"errors"
	"time"
)

// Transaction represents a transaction in the account
type Transaction struct {
	Amount      float64
	Date        time.Time
	Description string
}

// Account represents a bank account
type Account struct {
	ID           string
	Owner        string
	balance      float64
	transactions []Transaction
}

// NewAccount creates a new account with an initial balance
func NewAccount(id, owner string, initialBalance float64) *Account {
	// TODO: Implement this function
	return nil
}

// Deposit adds funds to the account
// Returns error if the amount is negative
func (a *Account) Deposit(amount float64, description string) error {
	// TODO: Implement this function
	return nil
}

// Withdraw removes funds from the account
// Returns error if the amount is negative or if there is insufficient balance
func (a *Account) Withdraw(amount float64, description string) error {
	// TODO: Implement this function
	return nil
}

// Balance returns the current account balance
func (a *Account) Balance() float64 {
	// TODO: Implement this function
	return 0
}

// Statement returns the transaction history
func (a *Account) Statement() []Transaction {
	// TODO: Implement this function
	return nil
}