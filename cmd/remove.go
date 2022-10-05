package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCommand)
}
  
var removeCommand = &cobra.Command{
	Use:   "remove",
	Short: "Removes List Item",
	Long:  `Removes a list item from your list using the items ID`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Removing...")
	},
}