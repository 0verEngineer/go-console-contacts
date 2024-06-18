package commands

import (
	"console-contacts/pkg/contacts"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var AddContactCmd = &cobra.Command{
	Use:   "add [contact book] [first name] [last name] [email] [phone number]",
	Short: "Add a new contact to the specified contact book",
	Args:  cobra.RangeArgs(5, 5),
	Run: func(cmd *cobra.Command, args []string) {
		anyInvalid := false
		if !nameValid(args[1]) {
			fmt.Println("First name is invalid")
			anyInvalid = true
		}
		if !nameValid(args[2]) {
			fmt.Println("Last name is invalid")
			anyInvalid = true
		}
		if !(emailValid(args[3]) && isStringValidForFile(args[2])) {
			fmt.Println("E-Mail is invalid")
			anyInvalid = true
		}
		if !isStringValidForFile(args[4]) {
			fmt.Println("Phone number is invalid")
			anyInvalid = true
		}

		if anyInvalid {
			return
		}

		if !contacts.DoesBookExist(args[0], true) {
			return
		}

		id, err := contacts.GetNextLineNumber(args[0])
		if err != nil || id == -1 {
			fmt.Println("Unable to get next line number for the contact id")
			return
		}

		contacts.WriteLineToBook(args[0], "\""+strconv.Itoa(id)+"\";"+"\""+args[1]+"\";\""+args[2]+"\";\""+args[3]+"\";\""+args[4]+"\"")
	},
}
