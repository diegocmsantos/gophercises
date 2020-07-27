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
		var result string
		for _, todo := range values {
			result += fmt.Sprintf("%s, %t\n", todo.Description, todo.Done)
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
