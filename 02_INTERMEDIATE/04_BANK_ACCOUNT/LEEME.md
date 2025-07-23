# Ejercicio 4: Sistema de Cuenta Bancaria

## Descripción
Implementa un sistema simple de cuenta bancaria utilizando structs y métodos en Go.

## Requisitos
1. Crea un struct `Account` en el archivo `account.go` con los siguientes campos:
   - `ID` (string)
   - `Owner` (string)
   - `Balance` (float64)
   - `transactions` (slice de Transaction, privado)

2. Crea un struct `Transaction` con los siguientes campos:
   - `Amount` (float64)
   - `Date` (time.Time)
   - `Description` (string)

3. Implementa los siguientes métodos para `Account`:
   - `NewAccount(id, owner string, initialBalance float64) *Account` - Constructor
   - `Deposit(amount float64, description string) error` - Añade fondos
   - `Withdraw(amount float64, description string) error` - Retira fondos
   - `Balance() float64` - Devuelve el saldo actual
   - `Statement() []Transaction` - Devuelve el historial de transacciones

4. Consideraciones:
   - No se permite retirar más dinero del disponible
   - No se permiten depósitos o retiros de cantidades negativas
   - Cada transacción debe registrarse con la fecha actual

## Pruebas
Ejecuta `go test` para verificar tu implementación.