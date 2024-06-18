package commands

import (
	"console-contacts/pkg/contacts"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"text/tabwriter"
)

var ListCmd = &cobra.Command{
	Use:   "list [contact book]",
	Short: "Lists all contacts in the given book",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if contacts.DoesBookExist(args[0], true) {
			listContacts(args[0])
		}
	},
}

func listContacts(book string) {
	allContacts, err := contacts.ReadAllContacts(book)
	if err != nil {
		fmt.Println("Unable to list all contacts")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	_, err = fmt.Fprintln(w, "ID\tFirst name\tLast name\tE-Mail\tPhone number")
	if err != nil {
		fmt.Println("Error printing list header: ", err)
		return
	}

	for i := range allContacts {
		_, err := fmt.Fprintln(w, strconv.Itoa(allContacts[i].Id)+"\t"+allContacts[i].FirstName+"\t"+allContacts[i].LastName+"\t"+allContacts[i].Email+"\t"+allContacts[i].PhoneNumber)
		if err != nil {
			fmt.Println("Error printing contact: ", err)
			return
		}
	}

	if err := w.Flush(); err != nil {
		fmt.Println("Error flushing printed list: ", err)
	}
}
