package main

import (
	"console-contacts/cmd/console-contacts/commands"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of the application.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("console-contacts 0.0.1")
		},
	}

	var rootCmd = &cobra.Command{Use: "console-contacts"}
	rootCmd.AddCommand(versionCmd, commands.CreateCmd, commands.ClearCmd, commands.DeleteCmd, commands.AddContactCmd, commands.DeleteContactCmd, commands.ListCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
