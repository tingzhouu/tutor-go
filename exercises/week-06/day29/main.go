package main

import (
	"fmt"
	"io"
	"os"
)

type Event struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}

func writeReport(w io.Writer, events []Event) {
	for _, e := range events {
		fmt.Fprintf(w, "id: %s; type: %s; amount: %d\n", e.ID, e.Type, e.Amount)
	}
}

func main() {
	events := []Event{
		{ID: "id-1", Type: "type-a", Amount: 100},
		{ID: "id-2", Type: "type-b", Amount: 200},
		{ID: "id-3", Type: "type-c", Amount: 300},
	}

	// use stdout
	writeReport(os.Stdout, events)

	// use file
	file, err := os.Create("report.txt")
	if err != nil {
		return
	}
	writeReport(file, events)
	myDoSomething(&MyFile{})
}

type MyWriter interface {
	MyWrite(p []byte) (n int, err error)
}

type MyFile struct {
}

func (m *MyFile) MyWrite(p []byte) (n int, err error) {
	return
}

func myDoSomething(m MyWriter) {

}
