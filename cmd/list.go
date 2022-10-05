package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCommand)
}
  
var listCommand = &cobra.Command{
	Use:   "remove",
	Short: "Prints TODO list",
	Long:  `Prints TODO list of all current list items.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing...")
	},
}