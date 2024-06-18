package commands

import (
	"console-contacts/pkg/contacts"
	"github.com/spf13/cobra"
)

var ClearCmd = &cobra.Command{
	Use:   "clear [book to clear]",
	Short: "Deletes all saved contacts from specified contact book",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		contacts.CreateOrClearBook(args[0], true)
	},
}
