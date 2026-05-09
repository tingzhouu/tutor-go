package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (u User) MarshalJSON() ([]byte, error) {
	res, err := json.Marshal(struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
	}{
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	})
	if err != nil {
		fmt.Printf("err: %v", err)
		return []byte(`{}`), nil
	}
	return res, nil
}

func (u *User) UnmarshalJSON(b []byte) error {
	val := struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
	}{}
	err := json.Unmarshal(b, &val)
	if err != nil {
		fmt.Printf("err: %v", err)
		return err
	}
	u.Name = val.Name + "UNMARSHALLED"
	u.Email = val.Email + "UNMARSHALLED"
	u.CreatedAt = val.CreatedAt + "UNMARSHALLED"
	return nil
}

func main() {
	e := json.NewEncoder(os.Stdout)
	err := e.Encode(User{Name: "John", Email: "john@email.com", CreatedAt: "2026-03-01"})
	if err != nil {
		fmt.Printf("error %v", err)
	}

	var user User
	r := strings.NewReader(`{"name":"John","email":"john@email.com","created_at":"2026-03-01"}`)
	d := json.NewDecoder(r)
	d.Decode(&user)
	fmt.Printf("Name: %s; Email: %s; Created At: %s\n", user.Name, user.Email, user.CreatedAt)
}
