package main

import (
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Printf("error occurred when opening file %v", err)
		panic(err)
	}
}

func safeDiv(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return a / b, nil
}

func main() {
	file, err := os.Open("exercises/week-09/day48/file")
	if err != nil {
		fmt.Printf("error occurred when opening file %v", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	check(err)
	fmt.Printf("%d bytes: %s\n", len(content), string(content))

	fmt.Println("read file - done")

	fmt.Println("start safeDiv")
	res, err := safeDiv(5, 0)
	fmt.Printf("res: %d\n err: %v\n\n", res, err)
	fmt.Println("end safeDiv")
}
