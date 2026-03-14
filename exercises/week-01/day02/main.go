package main

import "fmt"

func main() {
	// =============================================
	// 1. The three ways to declare variables in Go
	// =============================================

	// Way 1: var with explicit type (like `let x: number = 10` in TS)
	var age int = 30

	// Way 2: var with type inference (like `let x = 10`)
	var name = "Ting"

	// Way 3: short declaration with := (most common, only inside functions)
	city := "Singapore"

	fmt.Println(age, name, city)

	// =============================================
	// 2. Zero values — Go's answer to null/undefined
	// =============================================
	// In TS: let x: number; // undefined
	// In Go: every type has a zero value. No surprises.

	var zeroInt int
	var zeroStr string
	var zeroBool bool
	var zeroFloat float64

	fmt.Println("Zero int:", zeroInt)     // 0
	fmt.Println("Zero string:", zeroStr)  // "" (empty string)
	fmt.Println("Zero bool:", zeroBool)   // false
	fmt.Println("Zero float:", zeroFloat) // 0

	// =============================================
	// 3. No implicit coercion — this is different from TS!
	// =============================================
	// In TS: 5 + 3.2 = 8.2 (works fine)
	// In Go: you MUST convert explicitly

	whole := 5
	decimal := 3.2

	// This would NOT compile:
	// result := whole + decimal

	// You must convert:
	result := float64(whole) + decimal
	fmt.Println("5 + 3.2 =", result)

	// Same with int sizes — int32 + int64 won't compile either
	var a int32 = 100
	var b int64 = 200
	sum := int64(a) + b
	fmt.Println("100 + 200 =", sum)

	// =============================================
	// 4. Constants — like TS `const`, but truly immutable
	// =============================================
	// In TS: const obj = {x: 1}; obj.x = 2; // allowed!
	// In Go: const is for primitive values only, and they're truly constant

	const pi = 3.14159
	const greeting = "hello"
	// const can't be used with := syntax
	// const also can't be a struct or slice — only basic types

	fmt.Println(pi, greeting)

	// =============================================
	// EXERCISE: Your turn!
	// =============================================
	// Uncomment and fix the code below. Each snippet has a type error.
	// Think about WHY Go rejects each one.

	// --- Fix 1: Type mismatch ---
	var score float64 = 95.5

	// --- Fix 2: Implicit conversion ---
	var meters int = 100
	var kilometers float64 = float64(meters) / 1000

	// --- Fix 3: String + number concatenation ---
	var count int = 42
	var message string = fmt.Sprintf("Items: %d", count)

	fmt.Println("score", score)
	fmt.Println("kilometers", kilometers)
	fmt.Println("message", message)
}
