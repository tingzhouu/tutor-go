package main

import "fmt"

func filterAbove(orders []int, threshold int) []int {
	filteredRes := []int{}
	for _, price := range orders {
		if price >= threshold {
			filteredRes = append(filteredRes, price)
		}
	}
	return filteredRes
}

func applyTax(orders []int, taxPercent int) []int {
	result := []int{}

	for _, price := range orders {
		result = append(result, price*(100+taxPercent)/100)
	}
	return result
}

func cartTotal(amounts []int) int {
	total := 0
	for _, amount := range amounts {
		total += amount
	}
	return total
}

func main() {
	// ============================================================
	// PART 1: Go's "for" does everything
	// ============================================================

	// Classic for (same as JS/C)
	for i := 0; i < 3; i++ {
		fmt.Println("classic:", i)
	}

	// "while" — just drop the init and post
	count := 0
	for count < 3 {
		fmt.Println("while-style:", count)
		count++
	}

	// infinite loop — for {} (like while(true))
	// for {
	//     break // would run forever without this
	// }

	// ============================================================
	// PART 2: range over slices
	// JS:  items.forEach((item, i) => ...)
	// Go:  for i, item := range items { ... }
	// ============================================================

	prices := []int{1999, 2499, 599, 3999, 4999}

	// index + value
	for i, p := range prices {
		fmt.Printf("  [%d] %d cents\n", i, p)
	}

	// value only (discard index)
	for _, p := range prices {
		fmt.Printf("  %d cents\n", p)
	}

	// index only
	for i := range prices {
		fmt.Printf("  index %d\n", i)
	}

	// ============================================================
	// PART 3: range over maps
	// JS:  Object.entries(m).forEach(([k, v]) => ...)
	// Go:  for k, v := range m { ... }
	// NOTE: map iteration order is NOT guaranteed (unlike JS objects)
	// ============================================================

	cart := map[string]int{
		"shirt": 2499,
		"shoes": 8999,
		"hat":   1599,
	}

	for item, price := range cart {
		fmt.Printf("  %s: $%.2f\n", item, float64(price)/100)
	}

	// ============================================================
	// EXERCISES
	// Rewrite these JS patterns in Go. Each is a common
	// data transformation you'd do in a payment system.
	// ============================================================

	orders := []int{4999, 1299, 7500, 350, 2199}

	// -----------------------------------------------------------
	// EXERCISE 1: Filter
	// JS:  orders.filter(o => o >= 2000)
	// Return a new slice with only orders >= 2000 cents
	// -----------------------------------------------------------

	// TODO: implement filterAbove
	expensive := filterAbove(orders, 2000)
	fmt.Println("expensive orders:", expensive)

	// -----------------------------------------------------------
	// EXERCISE 2: Map (transform)
	// JS:  orders.map(o => o * 1.1)  // add 10% tax
	// Return a new slice with 10% added to each amount
	// (use integer math: amount * 110 / 100)
	// -----------------------------------------------------------

	// TODO: implement applyTax
	taxed := applyTax(orders, 10)
	fmt.Println("with 10% tax:", taxed)

	// -----------------------------------------------------------
	// EXERCISE 3: Reduce (already did this yesterday!)
	// But this time, combine filter + map + reduce:
	// "Total of orders >= 2000, after 10% tax"
	// Chain your functions together.
	// -----------------------------------------------------------

	// TODO: chain filterAbove → applyTax → cartTotal from yesterday
	result := cartTotal(applyTax(filterAbove(orders, 2000), 10))
	fmt.Printf("filtered + taxed total: %d cents ($%.2f)\n", result, float64(result)/100)
}

