package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := cobra.Command{
		Use: "task",
	}

	addTask := cobra.Command{
		Use: "add",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Printf("add received args %v\n", args)
			if len(args) < 1 {
				fmt.Printf("insufficient args\n")
				return
			}

			tasks := loadTasks()
			lastId := len(tasks)
			tasks = append(tasks, Task{
				ID:   lastId,
				Name: args[0],
				Done: false,
			})
			saveTasks(tasks)
		},
	}

	listTasks := cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("list received args %v\n", args)

			tasks := loadTasks()
			fmt.Printf("There are %d tasks\n", len(tasks))
			for _, t := range tasks {
				fmt.Printf("Task ID: %d; Name: %s; Done: %v\n", t.ID, t.Name, t.Done)
			}
		},
	}

	taskCompleted := cobra.Command{
		Use: "done",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("done received args %v\n", args)

			taskId, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("Received invalid task id %v", err)
				return
			}

			tasks := loadTasks()
			for i := range tasks {
				if tasks[i].ID == taskId {
					if tasks[i].Done {
						fmt.Printf("Task id %d already completed\n", tasks[i].ID)
						return
					}
					tasks[i].Done = true
					fmt.Printf("Updated task id %d as done\n", tasks[i].ID)
					saveTasks(tasks)
					return
				}
			}
			fmt.Printf("Received invalid task id %v\n", err)

		},
	}
	rootCmd.AddCommand(&addTask)
	rootCmd.AddCommand(&listTasks)
	rootCmd.AddCommand(&taskCompleted)
	rootCmd.Execute()
}
