package cmd

import (
	"fmt"
	"gophercises/todo/manager"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new todo to your list",
	Long:  "Adds a new todo to your list",
	Run: func(cmd *cobra.Command, args []string) {
		todo := manager.Todo{Description: args[0], Done: false}
		_, err := manager.Create(&todo)
		if err == nil {
			fmt.Fprintf(os.Stdin, "New todo %+v created", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
