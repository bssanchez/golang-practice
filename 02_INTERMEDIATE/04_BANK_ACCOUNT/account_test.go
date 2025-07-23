package account

import (
	"testing"
	"time"
)

func TestNewAccount(t *testing.T) {
	acc := NewAccount("1", "John Doe", 100.0)
	
	if acc.ID != "1" {
		t.Errorf("Expected ID to be '1', got %s", acc.ID)
	}
	
	if acc.Owner != "John Doe" {
		t.Errorf("Expected Owner to be 'John Doe', got %s", acc.Owner)
	}
	
	if acc.Balance() != 100.0 {
		t.Errorf("Expected Balance to be 100.0, got %f", acc.Balance())
	}
	
	if len(acc.Statement()) != 1 {
		t.Errorf("Expected 1 initial transaction, got %d", len(acc.Statement()))
	}
}

func TestDeposit(t *testing.T) {
	acc := NewAccount("1", "John Doe", 100.0)
	
	// Test valid deposit
	err := acc.Deposit(50.0, "Salary")
	if err != nil {
		t.Errorf("Unexpected error on valid deposit: %v", err)
	}
	
	if acc.Balance() != 150.0 {
		t.Errorf("Expected Balance to be 150.0 after deposit, got %f", acc.Balance())
	}
	
	// Test negative deposit
	err = acc.Deposit(-50.0, "Invalid")
	if err == nil {
		t.Error("Expected error on negative deposit, got nil")
	}
	
	// Balance should remain unchanged after failed deposit
	if acc.Balance() != 150.0 {
		t.Errorf("Expected Balance to remain 150.0 after failed deposit, got %f", acc.Balance())
	}
	
	// Check transaction history
	transactions := acc.Statement()
	if len(transactions) != 2 { // Initial deposit + salary
		t.Errorf("Expected 2 transactions, got %d", len(transactions))
	}
	
	lastTx := transactions[len(transactions)-1]
	if lastTx.Amount != 50.0 || lastTx.Description != "Salary" {
		t.Errorf("Last transaction incorrect: amount=%f, desc=%s", lastTx.Amount, lastTx.Description)
	}
}

func TestWithdraw(t *testing.T) {
	acc := NewAccount("1", "John Doe", 100.0)
	
	// Test valid withdrawal
	err := acc.Withdraw(30.0, "Groceries")
	if err != nil {
		t.Errorf("Unexpected error on valid withdrawal: %v", err)
	}
	
	if acc.Balance() != 70.0 {
		t.Errorf("Expected Balance to be 70.0 after withdrawal, got %f", acc.Balance())
	}
	
	// Test negative withdrawal
	err = acc.Withdraw(-30.0, "Invalid")
	if err == nil {
		t.Error("Expected error on negative withdrawal, got nil")
	}
	
	// Test excessive withdrawal
	err = acc.Withdraw(100.0, "Too much")
	if err == nil {
		t.Error("Expected error on excessive withdrawal, got nil")
	}
	
	// Balance should remain unchanged after failed withdrawal
	if acc.Balance() != 70.0 {
		t.Errorf("Expected Balance to remain 70.0 after failed withdrawal, got %f", acc.Balance())
	}
	
	// Check transaction history
	transactions := acc.Statement()
	if len(transactions) != 2 { // Initial deposit + groceries
		t.Errorf("Expected 2 transactions, got %d", len(transactions))
	}
	
	lastTx := transactions[len(transactions)-1]
	if lastTx.Amount != -30.0 || lastTx.Description != "Groceries" {
		t.Errorf("Last transaction incorrect: amount=%f, desc=%s", lastTx.Amount, lastTx.Description)
	}
}

func TestStatement(t *testing.T) {
	acc := NewAccount("1", "John Doe", 100.0)
	
	// Add some transactions
	acc.Deposit(50.0, "Salary")
	acc.Withdraw(30.0, "Groceries")
	acc.Deposit(10.0, "Refund")
	
	// Check transaction history
	transactions := acc.Statement()
	if len(transactions) != 4 { // Initial + 3 more
		t.Errorf("Expected 4 transactions, got %d", len(transactions))
	}
	
	// Check transaction order (should be chronological)
	for i := 1; i < len(transactions); i++ {
		if transactions[i].Date.Before(transactions[i-1].Date) {
			t.Error("Transactions not in chronological order")
		}
	}
	
	// Verify transaction amounts
	expectedAmounts := []float64{100.0, 50.0, -30.0, 10.0}
	for i, tx := range transactions {
		if tx.Amount != expectedAmounts[i] {
			t.Errorf("Transaction %d: expected amount %f, got %f", i, expectedAmounts[i], tx.Amount)
		}
	}
	
	// Verify final balance
	if acc.Balance() != 130.0 {
		t.Errorf("Expected final balance 130.0, got %f", acc.Balance())
	}
}

func TestTransactionDates(t *testing.T) {
	acc := NewAccount("1", "John Doe", 100.0)
	
	// Check that transaction dates are recent
	now := time.Now()
	transactions := acc.Statement()
	
	for _, tx := range transactions {
		// Transaction should be within the last second
		if now.Sub(tx.Date).Seconds() > 1 {
			t.Errorf("Transaction date too old: %v", tx.Date)
		}
	}
}