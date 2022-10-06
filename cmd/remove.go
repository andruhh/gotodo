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
		var currentItem TodoItem
		var confirmationText string

		if selectedNumber == 0 {
			fmt.Println("Please enter the number of the TODO item you wish to edit.")
			fmt.Scan(&selectedNumber)
		}

		for _, item := range todoList {
			if item.number == selectedNumber {
				fmt.Println("Type 'YES' to confirm removal")
				fmt.Scan(&confirmationText)

				if confirmationText != "YES" {
					fmt.Println("Incorrect confirmation. Aborting.")
					break
				}

				currentItem = item
				fmt.Printf("Removing %v from your list.\n", currentItem.text)
				removeItem(selectedNumber - 1)
				break
			}
		}
	},
}