package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	field   string
	message string
}

var ErrAccountFrozen = errors.New("Account Frozen")
var ErrInsufficientFunds = errors.New("Insufficient Funds")

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s - %s\n", v.field, v.message)
}

func withdrawFromAccount(accountID string, amount int) (string, error) {
	if amount <= 0 {
		return "", &ValidationError{field: "amount", message: "<=0"}
	}

	if accountID == "frozen" {
		return "", ErrAccountFrozen
	}

	if accountID == "" {
		return "", &ValidationError{field: "accountID", message: "required"}
	}

	if amount > 5000 {
		return "", fmt.Errorf("amount %d failed: %w", amount, ErrInsufficientFunds)
	}
	return "withdrawal successful", nil
}

func main() {
	var validationError *ValidationError
	res, err := withdrawFromAccount("account_1", 100)

	fmt.Printf("res %s\n", res)

	res, err = withdrawFromAccount("account_1", -1)

	if errors.As(err, &validationError) {
		fmt.Printf("validation error caught - %s\n", err.Error())
	}

	res, err = withdrawFromAccount("", 100)

	if errors.As(err, &validationError) {
		fmt.Printf("validation error caught - %s\n", err.Error())
	}

	res, err = withdrawFromAccount("frozen", 100)

	if errors.Is(err, ErrAccountFrozen) {
		fmt.Printf("frozen account error caught\n")
	}

	res, err = withdrawFromAccount("account_1", 10000)

	if errors.Is(err, ErrInsufficientFunds) {
		fmt.Printf("insufficient funds error caught\n")
	}

}
