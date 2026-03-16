package main

// This file imports from your payments package.
// It will NOT compile until you fix the visibility in payments/order.go.
//
// Once you've fixed everything, uncomment the code below and run:
//   go run exercises/week-02/day09/main.go

import (
	"fmt"

	"tutor-go/exercises/week-02/day09/payments"
)

func main() {
	// Uses the exported constructor
	order := payments.NewOrder("ord_456", 5999)
	fmt.Println(order.Summary())

	// Uses exported method
	order.Complete()
	fmt.Println(order.Summary())

	// Uses exported method (your Day 8 refund, now in a package)
	order.Refund()
	fmt.Println(order.Summary())

	// This should NOT compile — uncomment to verify:
	// order.addNote("hacking the internals")  // unexported method
	// order.internalNote = "oops"             // unexported field
	// fmt.Println(order.status)               // unexported field
}
