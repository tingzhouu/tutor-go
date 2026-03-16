package main

import "fmt"

func main() {
	// ============================================================
	// PART 1: Maps — make is REQUIRED before writing
	// ============================================================

	// var m map[string]int   // nil map — reads return zero, writes PANIC
	// m["key"] = 1           // panic: assignment to entry in nil map

	// Three ways to create a usable map:

	// Way 1: make (empty map, ready to use)
	m1 := make(map[string]int)
	m1["a"] = 1

	// Way 2: literal (if you know the initial values)
	m2 := map[string]int{"a": 1, "b": 2}

	// Way 3: make with size hint (optimization — NOT a max size)
	// "I expect about 100 entries" — avoids internal resizing
	m3 := make(map[string]int, 100)
	m3["a"] = 1 // can still add more than 100

	fmt.Println("m1:", m1, "m2:", m2, "m3 len:", len(m3))

	// TODO 1: predict — what does reading from a nil map do?
	var nilMap map[string]int
	val := nilMap["missing"]
	fmt.Println("nil map read:", val) // predict: ??? it will panic

	// ============================================================
	// PART 2: Slices — make lets you control length and capacity
	// ============================================================

	// Why would you make a slice instead of using []int{} ?
	// Answer: when you know the size upfront and want to avoid resizing.

	// make([]T, length, capacity)
	//   length   = how many elements exist right now (filled with zero values)
	//   capacity = how much space is allocated (optional)

	// Way 1: literal
	s1 := []int{1, 2, 3}

	// Way 2: make with length (pre-filled with zeros)
	s2 := make([]int, 5)

	// Way 3: make with length 0, capacity 5
	// "I'll append up to 5 items" — avoids reallocations
	s3 := make([]int, 0, 5)

	fmt.Println("s1:", s1, "len:", len(s1), "cap:", cap(s1))
	fmt.Println("s2:", s2, "len:", len(s2), "cap:", cap(s2))
	fmt.Println("s3:", s3, "len:", len(s3), "cap:", cap(s3))

	// TODO 2: predict — what happens here?
	s4 := make([]int, 3)
	s4 = append(s4, 99)
	fmt.Println("s4:", s4) // predict: ??? [99]

	// ============================================================
	// PART 3: make vs new vs literal — when to use what
	//
	// make(T)     → slices, maps, channels ONLY. Initializes internals.
	// new(T)      → allocates memory, returns *T. Rarely used in practice.
	// T{}         → struct/array literal. Most common for structs.
	// &T{}        → same as new but with initial values. Very common.
	// ============================================================

	// In practice, you'll mostly use:
	//   Structs:  p := Payment{Amount: 100}   or   p := &Payment{Amount: 100}
	//   Maps:     m := make(map[K]V)           or   m := map[K]V{"a": 1}
	//   Slices:   s := []int{1,2,3}            or   s := make([]int, 0, n)

	// ============================================================
	// EXERCISE: Build a frequency counter
	// Given a slice of strings, return a map of string → count.
	//
	// TS equivalent:
	//   const freq = (items: string[]): Record<string, number> =>
	//     items.reduce((acc, item) => ({ ...acc, [item]: (acc[item] || 0) + 1 }), {})
	//
	// Use make with a size hint since you know the max possible keys.
	// ============================================================

	items := []string{"shirt", "hat", "shirt", "shoes", "hat", "shirt"}

	// TODO 3: implement frequency and uncomment
	counts := frequency(items)
	fmt.Println("counts:", counts)
	// Expected: map[hat:2 shirt:3 shoes:1]
	_ = items
}

// TODO 3: implement frequency
func frequency(items []string) map[string]int {
	// frequencyMap := make(map[string]int)
	frequencyMap := map[string]int{}
	for _, item := range items {
		frequencyMap[item] += 1
	}
	return frequencyMap
}
