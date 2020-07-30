package cmd

import (
	"fmt"
	"gophercises/todo/manager"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a todo as done",
	Long:  "Marks a todo as done",
	Run: func(cmd *cobra.Command, args []string) {

		// getting all tasks
		todos, err := manager.List()
		if err != nil {
			fmt.Println("Error getting all tasks:", err)
		}

		// creating a list of integer ids
		var ids []int
		for _, arg := range args {
			id, _ := strconv.Atoi(arg)
			ids = append(ids, id)
		}

		// iterating over all ids
		for _, id := range ids {
			// getting the task for the chose id
			todo := todos[id-1]
			// deleting the found todo
			err := manager.MarkAsDone(todo.Key)
			if err != nil {
				fmt.Printf("Error marking \"%d\" as complete. Error: %q\n", todo.Key, err)
			} else {
				fmt.Printf("Todo \"%s\" marked as complete.\n", todo.Value)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
