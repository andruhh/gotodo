package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCommand)
}
  
var editCommand = &cobra.Command{
	Use:   "remove",
	Short: "Edits List Item",
	Long:  `Edits a list item from your list using the items ID`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Editing...")
	},
}