package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCommand)
}
  
var editCommand = &cobra.Command{
	Use:   "edit",
	Short: "Edits List Item",
	Long:  `Edits a list item from your list using the items ID`,
	Run: func(cmd *cobra.Command, args []string) {
		if selectedNumber == 0 {
			fmt.Println("Please enter the number of the ticket you wish to edit.")
			fmt.Scan(&selectedNumber)
		}

		for index, item := range todoList {
			if item.number == selectedNumber {
				var text string
				fmt.Printf("Please enter the text you wish to replace '%v' with:\n", item.text)
				fmt.Scan(&text)
				updateItemText(index, text)
				break
			}
		}
	},
}