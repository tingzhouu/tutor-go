package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s - %s", v.Field, v.Message)
}

type User struct {
	id int
}

var ErrNotFound = errors.New("user not found")

func findUser(id int) (User, error) {
	if id == -1 {
		return User{}, &ValidationError{Field: "id", Message: "must be positive"}
	}
	if id == 999 {
		return User{}, fmt.Errorf("database lookup: %w", ErrNotFound)
	}

	if id == 1 {
		return User{id: 1}, nil
	}

	return User{}, nil
}

func main() {
	_, err1 := findUser(-1)
	var ve *ValidationError
	if errors.Is(err1, &ve) {
		fmt.Println(ve.Error())
	}

	errors.As(err1, ve)

	_, err2 := findUser(999)
	if errors.Is(err2, ErrNotFound) {
		fmt.Println(err2.Error())
	}

}
