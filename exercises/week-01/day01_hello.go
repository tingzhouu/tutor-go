// // Week 1, Day 1 — Hello Go
// // Your first exercise: get comfortable with the basics
// //
// // Instructions:
// // 1. Run this file: go run main.go
// // 2. Complete the TODOs below
// // 3. Discuss with your tutor what surprised you compared to TypeScript

// package main

// import "fmt"

// func main() {
// 	// 1. Go has zero values — every type has a default
// 	// TODO: Declare these variables without assigning values, then print them
// 	// What do you expect each to print? (Predict before running!)
// 	var myInt int
// 	var myString string
// 	var myBool bool
// 	var myFloat float64

// 	fmt.Println("Zero values:")
// 	fmt.Println("int:", myInt)       // What do you expect?
// 	fmt.Println("string:", myString) // What do you expect?
// 	fmt.Println("bool:", myBool)     // What do you expect?
// 	fmt.Println("float:", myFloat)   // What do you expect?

// 	// 2. Short variable declaration — Go's version of `const x = ...`
// 	// The := operator declares AND assigns
// 	name := "Nasi"
// 	year := 2026

// 	fmt.Println("\nHello", name, "in", year)

// 	// 3. Multiple return values — this doesn't exist in TypeScript
// 	// TODO: Call the divide function below and handle both returns
// 	// result, err := divide(10, 3)
// 	// What should you do with err?

// 	// 4. TODO: Try calling divide(10, 0) — what happens?
// 	// result, err := divide(10, 0)
// 	// if err != nil {
// 	// 	fmt.Println("error", err)
// 	// 	return
// 	// }
// 	fmt.Println(greet("john"))
// }

// func greet(name string) string {
// 	return "hello " + name
// }

// // divide returns both the result and an error
// // In TypeScript, you'd throw an Error. In Go, you return it.
// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a / b, nil
// }
