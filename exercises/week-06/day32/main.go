package main

import "fmt"

type Base struct {
	ID        int
	CreatedAt string
}

func (b Base) Info() string {
	return fmt.Sprintf("ID: %d, created: %s", b.ID, b.CreatedAt)
}

type User struct {
	Base
	Email string
}

type Admin struct {
	User
	Role string
}

func (a Admin) Info() string {
	return fmt.Sprintf("ID: %d, created: %s, role: %s", a.ID, a.CreatedAt, a.Role)
}

func main() {
	a := Admin{
		User: User{
			Base:  Base{ID: 5, CreatedAt: "2026-01-10"},
			Email: "abc@gmcil.com",
		},
		Role: "Superuser",
	}

	fmt.Println(a.ID)

	fmt.Println(a.Info())
}
