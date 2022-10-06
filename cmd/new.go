package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCommand)
}

var newCommand = &cobra.Command{
	Use:   "new",
	Short: "Adds a new TODO Item",
	Long:  `Adds a new TODO Item to your TODO List`,
	Run: func(cmd *cobra.Command, args []string) {
		if newItemText == "" {
			fmt.Println("Please enter the text for your TODO list item:")
			fmt.Scan(&newItemText)
		}

		var item = TodoItem {
			number: len(todoList) + 1,
			text: newItemText,
		}

		todoList = append(todoList, item)
		fmt.Printf("You have added '%v' to your TODO List\n", newItemText)
		fmt.Println(todoList)
	},
}