package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Event struct {
	Name string
	Date string
}

func (e Event) MarshalJSON() ([]byte, error) {
	t, _ := time.Parse("2006-01-02", e.Date)
	return json.Marshal(struct {
		Name string `json:"name"`
		Date string `json:"date"`
	}{
		Name: e.Name,
		Date: t.Format("Jan 2, 2006"),
	})
}

func (e *Event) UnmarshalJSON(data []byte) error {
	var raw struct {
		Name string `json:"name"`
		Date string `json:"date"`
	}
	json.Unmarshal(data, &raw)
	t, _ := time.Parse("Jan 2, 2006", raw.Date)
	e.Date = t.Format("2006-01-02")
	e.Name = raw.Name
	return nil
}

func main() {
	encoder := json.NewEncoder(os.Stdout)

	err := encoder.Encode(Event{"James", "2026-01-02"})
	if err != nil {
		fmt.Println("error occurred while encoding")
		return
	}

	r := strings.NewReader(`{"name":"Concert","date":"Jun 1, 2026"}`)
	decoder := json.NewDecoder(r)
	var e Event
	err = decoder.Decode(&e)
	if err != nil {
		fmt.Println("error occurred while decoding")
		return
	}
	fmt.Println(e)
}
