package main

import "fmt"

type Sizer interface {
	Size() int
}

type Namer interface {
	Name() string
}

type Describable interface {
	Sizer
	Namer
}

type File struct {
	name string
	size int
}

func (f File) Name() string {
	return f.name
}

func (f File) Size() int {
	return f.size
}

func printInfo(d Describable) {
	fmt.Printf("%s; %d\n", d.Name(), d.Size())
}

func main() {
	file := File{
		"myfile",
		20,
	}
	printInfo(file)
}
