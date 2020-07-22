package cmd

import (
	"fmt"
	"gophercises/todo/manager"
	"os"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a todo as done",
	Long:  "Marks a todo as done",
	Run: func(cmd *cobra.Command, args []string) {
		todo := manager.Todo{Description: args[0], Done: true}
		err := manager.MarkAsDone(&todo)
		if err == nil {
			fmt.Fprintf(os.Stdin, "%s marked as done", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
