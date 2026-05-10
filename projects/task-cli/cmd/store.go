package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func loadTasks() []Task {
	f, err := os.Open("tasks.json")
	if err != nil {
		fmt.Printf("could not open file %v\n", err)
		return nil
	}
	defer f.Close()

	var res []Task
	d := json.NewDecoder(f)
	d.Decode(&res)

	return res
}

func saveTasks(tasks []Task) {
	f, err := os.Create("tasks.json")
	if err != nil {
		fmt.Printf("could not create file %v\n", err)
		return
	}
	defer f.Close()

	e := json.NewEncoder(f)
	e.Encode(tasks)
}
