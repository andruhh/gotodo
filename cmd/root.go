package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type TodoItem struct {
	number int
	text string
}

var todoList = []TodoItem{{number: 1, text: "Get milk"}, {number:2, text:"feed dogs"}}
var selectedNumber int
var newItemText string

var rootCmd = &cobra.Command{
	Use:   "gotodo",
	Short: "TODO List CLI",
	Long: `Welcome to your TODO List. You can View, Add, Edit, or Remove.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	removeCommand.Flags().IntVar(&selectedNumber, "number", 0, "number of the item you wish to remove")
	editCommand.Flags().IntVar(&selectedNumber, "number", 0, "number of the item you wish to edit")
	newCommand.Flags().StringVarP(&newItemText, "text", "t", "", "Text of the list item")
}

func removeItem(index int) {
	list := append(todoList[:index], todoList[index + 1:]...)
	for index, item := range list {
		item.number = index + 1
		list[index] = item
	}
	todoList = list
}

func updateItemText(index int, text string) {
	item := todoList[index]
	item.text = text
	todoList[index] = item
}

