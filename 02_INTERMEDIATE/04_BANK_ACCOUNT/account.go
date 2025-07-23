package account

import (
	"errors"
	"time"
)

// Transaction representa una transacción en la cuenta
type Transaction struct {
	Amount      float64
	Date        time.Time
	Description string
}

// Account representa una cuenta bancaria
type Account struct {
	ID           string
	Owner        string
	balance      float64
	transactions []Transaction
}

// NewAccount crea una nueva cuenta con un saldo inicial
func NewAccount(id, owner string, initialBalance float64) *Account {
	// TODO: Implementar esta función
	return nil
}

// Deposit añade fondos a la cuenta
// Devuelve error si el monto es negativo
func (a *Account) Deposit(amount float64, description string) error {
	// TODO: Implementar esta función
	return nil
}

// Withdraw retira fondos de la cuenta
// Devuelve error si el monto es negativo o si no hay suficiente saldo
func (a *Account) Withdraw(amount float64, description string) error {
	// TODO: Implementar esta función
	return nil
}

// Balance devuelve el saldo actual de la cuenta
func (a *Account) Balance() float64 {
	// TODO: Implementar esta función
	return 0
}

// Statement devuelve el historial de transacciones
func (a *Account) Statement() []Transaction {
	// TODO: Implementar esta función
	return nil
}