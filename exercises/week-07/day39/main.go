package main

import "fmt"

func filter[T any](items []T, fn func(item T) bool) []T {
	res := make([]T, 0)
	for _, item := range items {
		if fn(item) {
			res = append(res, item)
		}
	}
	return res
}

func mapSlice[T any, R any](items []T, fn func(T) R) []R {
	var res []R
	for _, item := range items {
		res = append(res, fn(item))
	}
	return res
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbers = filter(numbers, func(num int) bool {
		return num%2 == 0
	})
	fmt.Println(numbers)

	items := []string{"hello", "hi", "world"}
	res := mapSlice(items, func(item string) int {
		return len(item)
	})

	fmt.Println(res)
}
