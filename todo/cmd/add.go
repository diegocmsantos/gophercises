package cmd

import (
	"fmt"
	"gophercises/todo/manager"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new todo to your list",
	Long:  "Adds a new todo to your list",
	Run: func(cmd *cobra.Command, args []string) {

		var errorsArray, successArray []string
		for _, arg := range args {
			todo := manager.Todo{Value: arg}
			_, err := manager.Create(&todo)
			if err != nil {
				errorsArray = append(errorsArray, fmt.Sprintf("Error creating %s task", arg))
			} else {
				successArray = append(successArray, fmt.Sprintf("New task \"%s\" created", arg))
			}
		}

		if errorsArray != nil {
			fmt.Fprint(os.Stdout, strings.Join(errorsArray, "\n"))
		}
		if successArray != nil {
			fmt.Fprint(os.Stdout, strings.Join(successArray, "\n"))
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
