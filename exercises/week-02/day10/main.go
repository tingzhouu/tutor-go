package main

import "fmt"

func main() {
	fmt.Println("=== BUG 1: The Disappearing Update ===")
	bug1()

	fmt.Println("\n=== BUG 2: The Infinite Cart ===")
	bug2()

	fmt.Println("\n=== BUG 3: The Missing Item ===")
	bug3()
}

// ============================================================
// BUG 1: The Disappearing Update
// A merchant updates a payment status, but the change vanishes.
// Find the bug and explain why it happens.
// ============================================================

type Payment struct {
	ID     string
	Amount int
	Status string
}

// func (p Payment) MarkPaid() {
func (p *Payment) MarkPaid() {
	p.Status = "paid"
}

func bug1() {
	payment := Payment{ID: "pay_001", Amount: 4999, Status: "pending"}
	payment.MarkPaid()
	fmt.Println("Status:", payment.Status) // Expected: "paid"
}

// ============================================================
// BUG 2: The Infinite Cart
// This function should return items over $20, but something
// goes wrong. Find the bug.
// ============================================================

func filterExpensive(items []int) []int {
	var result []int
	for i := 0; i < len(items); i++ {
		if items[i] >= 2000 {
			// items = append(items, items[i]) // should remove?? Not sure if there's another purpose
			result = append(result, items[i])
		}
	}
	return result
}

func bug2() {
	cart := []int{1500, 3999, 999, 2499, 500}
	// WARNING: this will loop forever as written — don't run until you find the bug!
	expensive := filterExpensive(cart)
	fmt.Println("Expensive items:", expensive)
	fmt.Println("Cart:", cart)
	fmt.Println("(uncomment after fixing the bug)")
}

// ============================================================
// BUG 3: The Missing Item
// We build a map of order IDs → amounts, but one item is
// always missing. Find the bug.
// ============================================================

type Order struct {
	ID     string
	Amount int
}

func buildOrderMap(orders []Order) map[string]int {
	var m map[string]int = make(map[string]int)
	for _, o := range orders {
		m[o.ID] = o.Amount
	}
	return m
}

func bug3() {
	orders := []Order{
		{ID: "ord_1", Amount: 1999},
		{ID: "ord_2", Amount: 4500},
		{ID: "ord_3", Amount: 750},
	}
	// WARNING: this will panic as written — don't run until you find the bug!
	result := buildOrderMap(orders)
	fmt.Println("Order map:", result)
	fmt.Println("(uncomment after fixing the bug)")
}
