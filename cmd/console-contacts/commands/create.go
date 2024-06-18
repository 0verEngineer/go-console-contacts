package commands

import (
	"console-contacts/pkg/contacts"
	"fmt"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create [book to create]",
	Short: "Creates a new contact book",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !contacts.DoesBookExist(args[0], false) {
			contacts.CreateOrClearBook(args[0], true)
		} else {
			fmt.Println("Book does already exist")
		}
	},
}
