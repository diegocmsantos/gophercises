package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	userLicense string
	rootCmd     = &cobra.Command{
		Use:   "task",
		Short: "Another TODO application",
		Long:  `You can create tasks, and mark as done your own tasks`,
	}
)

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.PersistentFlags().StringP("author", "a", "Diego Maia", "author name for copyright attribution")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "Diego Maia <diegocmsantos@gmail.com")
	viper.SetDefault("license", "apache")

}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
}
