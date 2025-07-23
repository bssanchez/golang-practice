# Exercise 4: Bank Account System

## Description
Implement a simple bank account system using structs and methods in Go.

## Requirements
1. Create an `Account` struct in the `account.go` file with the following fields:
   - `ID` (string)
   - `Owner` (string)
   - `Balance` (float64)
   - `transactions` (slice of Transaction, private)

2. Create a `Transaction` struct with the following fields:
   - `Amount` (float64)
   - `Date` (time.Time)
   - `Description` (string)

3. Implement the following methods for `Account`:
   - `NewAccount(id, owner string, initialBalance float64) *Account` - Constructor
   - `Deposit(amount float64, description string) error` - Add funds
   - `Withdraw(amount float64, description string) error` - Withdraw funds
   - `Balance() float64` - Return current balance
   - `Statement() []Transaction` - Return transaction history

4. Considerations:
   - Withdrawing more money than available is not allowed
   - Negative deposits or withdrawals are not allowed
   - Each transaction must be recorded with the current date

## Tests
Run `go test` to verify your implementation.