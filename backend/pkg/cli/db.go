package cli

import (
	"linker/internal/database"
	"log"

	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database commands for Linker",
}

func init() {
	_, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(dbCmd)
}
