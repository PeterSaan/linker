package cli

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clinker",
	Short: "Linker CLI",
}

var hideHelp = &cobra.Command{
	Hidden: true,
}

func Execute() {
	rootCmd.SetHelpCommand(hideHelp)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
