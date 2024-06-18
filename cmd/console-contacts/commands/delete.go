package commands

import (
	"console-contacts/pkg/contacts"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete-book [book to delete]",
	Short: "Deletes the specified contact book",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		exPath := contacts.GetExecutablePath()
		err := os.Remove(exPath + args[0] + ".txt")
		if err != nil {
			log.Fatal(err)
		}
	},
}

var DeleteContactCmd = &cobra.Command{
	Use:   "delete [contact book] [contact ID]",
	Short: "Deletes the specified contact from the specified contact book",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		book := args[0]
		allContacts, err := contacts.ReadAllContacts(book)
		if err != nil {
			fmt.Println("Unable to read contacts from book", book)
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("ID invalid")
			return
		}

		allContacts = contacts.RemoveEntryByIdAndChangeIdsAccordingly(allContacts, id)
		contacts.CreateOrClearBook(book, false)
		contacts.WriteAllContacts(book, allContacts)
	},
}
