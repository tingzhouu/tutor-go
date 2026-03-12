package main

import "fmt"

// =============================================
// 1. Structs — Go's replacement for classes
// =============================================

// No constructor, no `this`, no visibility modifiers.
// Just fields. Uppercase = exported, lowercase = unexported.

type Payment struct {
	ID       string
	Amount   int // in cents
	Currency string
}

// Methods are functions with a "receiver" — the (p Payment) part.
// Think of it as: "this function belongs to Payment"

func (p Payment) AmountInDollars() float64 {
	return float64(p.Amount) / 100
}

func (p Payment) Summary() string {
	return fmt.Sprintf("%s: $%.2f %s", p.ID, p.AmountInDollars(), p.Currency)
}

// =============================================
// 2. Interfaces — implicit satisfaction
// =============================================

// In TS: class Payment implements Summarizable { ... }
// In Go: if Payment has a Summary() method, it satisfies the interface. Done.
// No "implements" keyword. This is called "structural typing" (TS has it too,
// but Go applies it to methods, not just shape).

type Summarizable interface {
	Summary() string
}

// This function accepts ANY type that has a Summary() method.
// Payment satisfies it without declaring anything.
func printSummary(s Summarizable) {
	fmt.Println(s.Summary())
}

// =============================================
// 3. Another type that satisfies the same interface
// =============================================

type Refund struct {
	ID        string
	PaymentID string
	Amount    int
}

func (r Refund) Summary() string {
	return fmt.Sprintf("%s: refund of %d cents on %s", r.ID, r.Amount, r.PaymentID)
}

type Charge struct {
	ID            string
	Amount        int
	CustomerEmail string
}

func (c Charge) Summary() string {
	return fmt.Sprintf("Charge %s for %d cents to email %s", c.ID, c.Amount, c.CustomerEmail)
}

// =============================================
// READING EXERCISE: Predict the output, then run it
// =============================================

func main() {
	p := Payment{
		ID:       "pay_123",
		Amount:   4999,
		Currency: "USD",
	}

	r := Refund{
		ID:        "ref_456",
		PaymentID: "pay_123",
		Amount:    1500,
	}

	// What will these print?
	fmt.Println(p.AmountInDollars()) // 44.99
	fmt.Println(p.Summary())         // pay_123 $44.99 USD
	fmt.Println(r.Summary())         // ref_456 refund of 1500 cents on pay_123

	// Both Payment and Refund satisfy Summarizable — no "implements" needed
	printSummary(p) // pay_123 $44.99 USD
	printSummary(r) // ref_456 refund of 1500 cents on pay_123

	// =============================================
	// EXERCISE: Your turn!
	// =============================================
	// 1. Create a "Charge" struct with fields: ID (string), Amount (int),
	//    CustomerEmail (string)
	c := Charge{
		ID:            "charge_789",
		Amount:        100,
		CustomerEmail: "abc@gmail.com",
	}
	printSummary(c)

	// 2. Add a Summary() method so it satisfies Summarizable
	// 3. Call printSummary with your Charge
}
