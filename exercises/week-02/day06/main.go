package main

import "fmt"

func main() {
	// ============================================================
	// PART 1: Arrays vs Slices
	// In JS, all arrays are dynamic. Go splits the concept in two.
	// ============================================================

	// Arrays: fixed size, value type (copied on assignment!)
	a := [3]int{10, 20, 30}
	b := a // What happens here? In JS, b would reference the same array.
	b[0] = 999

	fmt.Println("a:", a) // TODO 1: Predict the output before running // a: [10,20,30] b: [999,20,30]
	fmt.Println("b:", b)

	// Slices: dynamic, reference type (like JS arrays)
	s1 := []int{10, 20, 30}
	s2 := s1 // What about this one?
	s2[0] = 999

	fmt.Println("s1:", s1) // TODO 2: Predict the output before running // s1: [999,20,30] s2: [999,20,30]
	fmt.Println("s2:", s2)

	// ============================================================
	// PART 2: Slice operations
	// JS:  arr.push(4)       → Go: append(slice, 4)
	// JS:  arr.slice(1, 3)   → Go: slice[1:3]
	// JS:  arr.length        → Go: len(slice)
	// ============================================================

	prices := []int{1999, 2499, 599, 3999}

	// append returns a NEW slice (doesn't mutate in place like JS push)
	prices = append(prices, 4999)
	fmt.Println("after append:", prices)

	// Slicing works like JS slice — [start:end), end is exclusive
	mid := prices[1:3]
	fmt.Println("mid:", mid)

	// But watch out — mid shares the underlying array with prices!
	mid[0] = 0
	fmt.Println("prices after mid mutation:", prices) // TODO 3: Predict this // 1999, 0, 599, 3999, 4999

	// ============================================================
	// PART 3: Maps
	// JS:  const m = {}          → Go: m := make(map[string]int)
	// JS:  m["key"] = val        → Go: m["key"] = val  (same!)
	// JS:  delete m.key          → Go: delete(m, "key")
	// JS:  "key" in m            → Go: _, ok := m["key"]
	// ============================================================

	// Maps must be initialized — a nil map will panic on write
	cart := make(map[string]int)
	cart["shirt"] = 2499
	cart["shoes"] = 8999
	cart["hat"] = 1599

	fmt.Println("cart:", cart)
	fmt.Println("shirt price:", cart["shirt"])

	// The "comma ok" idiom — like checking if a key exists
	price, ok := cart["socks"]
	fmt.Println("socks price:", price, "found:", ok) // TODO 4: Predict this // 0, false

	// delete works like JS delete
	delete(cart, "hat")
	fmt.Println("after delete:", cart)

	// ============================================================
	// EXERCISE: Build a simple cart total function
	// Given a map of item→price (cents), return the total.
	// This is your parseAmount equivalent but for collections.
	//
	// In TypeScript you'd write:
	//   const cartTotal = (items: Record<string, number>): number =>
	//     Object.values(items).reduce((sum, p) => sum + p, 0)
	//
	// Write the Go version below.
	// ============================================================

	// TODO 5: Implement cartTotal and uncomment the lines below
	total := cartTotal(cart)
	fmt.Println("cart total:", total, "cents")
}

// TODO 5: Write cartTotal here
// func cartTotal(items map[string]int) int {
//     ...
// }

func cartTotal(items map[string]int) int {
	total := 0
	for _, price := range items {
		total += price
	}
	return total
}
