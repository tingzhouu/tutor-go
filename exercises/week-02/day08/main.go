package main

import "fmt"

type Payment struct {
	Amount int
	Status string
}

func main() {
	// ============================================================
	// PART 1: Pass by value (Go's default)
	// In JS, objects are passed by reference. In Go, structs are COPIED.
	// ============================================================

	original := Payment{Amount: 4999, Status: "pending"}
	copy := original // full copy, like array assignment in Go
	copy.Status = "completed"

	fmt.Println("original:", original.Status) // TODO 1: predict // Original: pending. copy: completed
	fmt.Println("copy:", copy.Status)

	// ============================================================
	// PART 2: Pointers — opting into references
	// & = "address of"    * = "value at address"
	// ============================================================

	x := 42
	p := &x // p is a *int (pointer to int)
	fmt.Println("x:", x)
	fmt.Println("p:", p)   // prints a memory address
	fmt.Println("*p:", *p) // dereference — gets the value

	*p = 100                            // modify x through the pointer
	fmt.Println("x after *p = 100:", x) // TODO 2: predict // x is 100

	// ============================================================
	// PART 3: Pointers with functions
	// This is the real reason pointers matter:
	// "Should this function modify the original, or work on a copy?"
	// ============================================================

	// This function CANNOT modify the original (receives a copy)
	// Think: TS primitive — changing a number param doesn't affect the caller
	payment := Payment{Amount: 4999, Status: "pending"}
	tryComplete(payment)
	fmt.Println("after tryComplete:", payment.Status) // TODO 3: predict // payment.Status is pending

	// This function CAN modify the original (receives a pointer)
	// Think: TS object reference — changes are visible to the caller
	completePayment(&payment)
	fmt.Println("after completePayment:", payment.Status) // TODO 4: predict // payment.Status is completed

	// ============================================================
	// PART 4: Method receivers — value vs pointer
	// This is where it really matters in day-to-day Go.
	// Yesterday's structs used value receivers. Now let's compare.
	// ============================================================

	order := Order{ID: "ord_123", Amount: 7500, Status: "pending"}

	order.Summary()                                // value receiver — can read but can't modify
	order.Complete()                               // pointer receiver — can modify
	fmt.Println("after Complete():", order.Status) // TODO 5: predict // order.status is completed

	// ============================================================
	// EXERCISE: Refund an order
	// Write a method on *Order that:
	// 1. Checks if status is "completed" (if not, print an error and return)
	// 2. Sets status to "refunded"
	// 3. Prints "Refunded order <ID> for <Amount> cents"
	//
	// In TS you'd write:
	//   class Order {
	//     refund() {
	//       if (this.status !== 'completed') throw new Error('...')
	//       this.status = 'refunded'
	//     }
	//   }
	//
	// In Go, "this" is the receiver. Use a pointer receiver to mutate.
	// ============================================================

	// TODO 6: implement Refund() and uncomment these lines
	order.Refund()
	fmt.Println("after Refund():", order.Status)
	//
	// // Try refunding again — should print an error
	order.Refund()
}

// Cannot modify original — receives a copy
func tryComplete(p Payment) {
	p.Status = "completed"
}

// CAN modify original — receives a pointer
func completePayment(p *Payment) {
	p.Status = "completed"
}

// Order type for method receiver examples
type Order struct {
	ID     string
	Amount int
	Status string
}

// Value receiver — gets a copy, can't modify the original
func (o Order) Summary() {
	fmt.Printf("Order %s: %d cents (%s)\n", o.ID, o.Amount, o.Status)
}

// Pointer receiver — gets the original, CAN modify it
func (o *Order) Complete() {
	o.Status = "completed"
}

// TODO 6: implement Refund on *Order
func (o *Order) Refund() {
	if o.Status != "completed" {
		fmt.Printf("Order id %s in status %s cannot be refunded\n", o.ID, o.Status)
		return
	}
	o.Status = "refunded"
	fmt.Printf("Refunded order %s for %d cents\n", o.ID, o.Amount)
}
