// Package payments provides types for processing payment orders.
//
// EXERCISE: Fix the visibility issues in this file.
// Some fields and methods are unexported (lowercase) when they
// should be exported, and vice versa.
//
// Rules:
//   - Types used by other packages must be Uppercase
//   - Fields that callers need to read must be Uppercase
//   - Fields that are internal implementation details stay lowercase
//   - Methods called from outside must be Uppercase
//   - Helper methods stay lowercase

package payments

import "fmt"

// TODO 1: This struct has visibility problems. Fix them.
// External code needs to: create orders, read ID/Amount/Status
// External code should NOT be able to: directly set the internalNote
type Order struct {
	ID           string
	Amount       int
	Status       string
	internalNote string
}

// TODO 2: Fix the constructor. In Go, NewXxx() is the convention
// for constructors (like a factory function).
// This needs to be callable from other packages.
func NewOrder(id string, amount int) Order {
	return Order{
		ID:           id,
		Amount:       amount,
		Status:       "pending",
		internalNote: "created via NewOrder",
	}
}

// TODO 3: Fix method visibility.
// Summary should be callable from outside.
// addNote is an internal helper — should stay unexported.
func (o Order) Summary() string {
	return fmt.Sprintf("Order %s: %d cents (%s)", o.ID, o.Amount, o.Status)
}

func (o *Order) addNote(note string) {
	o.internalNote = note
}

// TODO 4: Fix this method — it should be exported.
// It also needs to use the unexported addNote helper.
func (o *Order) Complete() {
	o.Status = "completed"
	o.addNote("completed by merchant")
}

// TODO 5: Add a Refund() method (exported, pointer receiver).
// Reuse your logic from Day 8, but adapt it to the new field names.
// Use addNote to record "refunded by merchant".

func (o *Order) Refund() {
	if o.Status != "completed" {
		fmt.Printf("Order %s for %d in status %s that cannot be refunded", o.ID, o.Amount, o.Status)
		return
	}
	o.addNote("refunded by merchant")
	o.Status = "refunded"
}
