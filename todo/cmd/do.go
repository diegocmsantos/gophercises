package cmd

import (
	"fmt"
	"gophercises/todo/manager"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a todo as done",
	Long:  "Marks a todo as done",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			key, _ := strconv.Atoi(arg)
			err := manager.MarkAsDone(key)
			if err == nil {
				fmt.Fprintf(os.Stdin, "%s marked as done", args[0])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
