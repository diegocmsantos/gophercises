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
			fmt.Println("add cmd: error reading all keys: %q", err)
		}
		fmt.Print("This should show all todos", values)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
