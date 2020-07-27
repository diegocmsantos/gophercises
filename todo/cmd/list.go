package cmd

import (
	"fmt"
	"gophercises/todo/manager"
	"os"

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
		result := "You have the following tasks:\n"
		for i, todo := range values {
			result += fmt.Sprintf("%d. %s\n", i+1, todo.Description)
		}
		if result != "" {
			fmt.Fprint(os.Stdout, result)
		} else {
			fmt.Fprint(os.Stdout, "No tasks found.")
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
