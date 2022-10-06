package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCommand)
}
  
var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints TODO list",
	Long:  `Prints TODO list of all current TODO items.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(todoList) <= 0 {
			fmt.Println("No items in list.")
		} else {
			for _, item := range todoList {
				fmt.Printf("%v. %v\n", item.number, item.text)
			}
		}
	},
}