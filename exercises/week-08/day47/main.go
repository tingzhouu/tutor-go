package main

import "fmt"

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func middleware(name string, fn func()) {
	fmt.Printf("[%s] start\n", name)
	fn()
	fmt.Printf("[%s] end\n", name)
}

func compose(fns ...func()) func() {
	return func() {
		for _, fn := range fns {
			fn()
		}
	}
}

func main() {
	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	add5 := makeAdder(5)
	fmt.Println(add5(3))  // 8
	fmt.Println(add5(10)) // 15

	hello := func() { fmt.Println("hello") }
	bye := func() { fmt.Println("bye") }

	both := compose(
		func() { middleware("greet", hello) },
		func() { middleware("farewell", bye) },
	)
	both() // prints hello then bye
}
