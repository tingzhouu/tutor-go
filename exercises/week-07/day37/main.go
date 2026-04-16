package main

import (
	"errors"
	"fmt"
	"slices"
)

type ValidationError struct {
	field   string
	message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s - %s\n", v.field, v.message)
}

var ErrLimitExceeded = errors.New("Limit Exceeded")

var ErrUnsupportedCurrency = errors.New("unsupported currency")

func processPayment(amount int, currency string) (string, error) {
	if amount <= 0 {
		return "", &ValidationError{field: "amount", message: "Must be > 0"}
	}

	validCurrencies := []string{"USD", "SGD", "EUR"}

	if !slices.Contains(validCurrencies, currency) {
		return "", ErrUnsupportedCurrency
	}

	if amount > 10000 {
		return "", fmt.Errorf("processing payment of %d: %w %w", amount, ErrLimitExceeded)
	}

	return "payment processed", nil
}

func main() {
	_, err := processPayment(-10, "SGD")
	var valError *ValidationError
	if errors.As(err, &valError) {
		fmt.Println(valError.field)
	}

	_, err = processPayment(100, "JPY")
	if errors.Is(err, ErrUnsupportedCurrency) {
		fmt.Printf("%v\n", err.Error())
	}

	_, err = processPayment(50000, "USD")
	if errors.Is(err, ErrLimitExceeded) {
		fmt.Printf("exceeded limit err: %v\n", err.Error())
	}

	res, err := processPayment(200, "SGD")
	if err == nil {
		fmt.Println(res)
	}

}
