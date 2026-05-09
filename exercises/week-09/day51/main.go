package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (u User) MarshalJSON() ([]byte, error) {
	// u.CreatedAt is in 2026-12-25 format
	// we want to return the response as Jan 2, 2026
	// therefore what we need to do is to first parse the time, then format it

	t, _ := time.Parse("2006-01-02", u.CreatedAt)

	res, err := json.Marshal(struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
	}{
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: t.Format("Jan 2, 2006"),
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

	// created_at is in format Jan 2, 2006. We want to store it as 2006-01-02
	t, _ := time.Parse("Jan 2, 2006", val.CreatedAt)

	u.Name = val.Name
	u.Email = val.Email
	u.CreatedAt = t.Format("2006-01-02")
	return nil
}

func main() {
	e := json.NewEncoder(os.Stdout)
	err := e.Encode(User{Name: "John", Email: "john@email.com", CreatedAt: "2026-03-01"})
	if err != nil {
		fmt.Printf("error %v", err)
	}

	var user User
	r := strings.NewReader(`{"name":"John","email":"john@email.com","created_at":"Mar 1, 2026"}`)
	d := json.NewDecoder(r)
	d.Decode(&user)
	fmt.Printf("Name: %s; Email: %s; Created At: %s\n", user.Name, user.Email, user.CreatedAt)
}
