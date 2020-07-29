package cmd

import (
	"fmt"
	"gophercises/todo/manager"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all existing todo's",
	Long:  "Lists all existing todo's",
	Run: func(cmd *cobra.Command, args []string) {

		values, err := manager.List()
		if err != nil {
			fmt.Println("add cmd: error reading all keys: %v", err)
		}
		if len(values) == 0 {
			fmt.Println("No tasks found!")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, todo := range values {
			fmt.Printf("%d. %s\n", i+1, todo.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
