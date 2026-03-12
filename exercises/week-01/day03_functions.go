// package main

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// 	"strconv"
// 	"strings"
// )

// // parseAmount takes a string like "$49.99" and returns the amount in cents.
// // In Go, instead of throwing an error, we return it as a second value.
// //
// // Your job: fill in the function body.
// // Hints:
// //   - strings.HasPrefix(s, "$") instead of startsWith
// //   - strconv.ParseFloat(s, 64) returns (float64, error)
// //   - You'll need to add "strings" and "strconv" to your imports
// //   - Use the "happy path on the left" idiom from Day 1
// //   - math.Round(f) rounds a float64

// func parseAmount(input string) (int, error) {
// 	// YOUR CODE HERE
// 	if !strings.HasPrefix(input, "$") {
// 		return 0, errors.New("does not have prefix $")
// 	}

// 	inputString := input[1:]
// 	inputFloat, err := strconv.ParseFloat(inputString, 64)
// 	if err != nil {
// 		return 0, fmt.Errorf("invalid format, not a number: %w", err)
// 	}

// 	return int(math.Round(inputFloat * 100)), nil
// }

// func main() {
// 	// Test cases — try all of these
// 	tests := []string{"$49.99", "$0.01", "49.99", "$abc", "$100"}

// 	for _, t := range tests {
// 		cents, err := parseAmount(t)
// 		if err != nil {
// 			fmt.Printf("parseAmount(%q) → ERROR: %s\n", t, err)
// 			continue
// 		}
// 		fmt.Printf("parseAmount(%q) → %d cents\n", t, cents)
// 	}
// }
