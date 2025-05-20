package cli

import (
	"linker/internal/database"
	"linker/internal/migration"
	"log"

	"github.com/spf13/cobra"
)

var dbMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Generate database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		println("Connecting to database")
		db := database.GetDB()

		println("Migrating...")
		if err := migration.Migrate(db); err != nil {
			log.Fatal(err)
		}

		println("Migration finished")
	},
}

func init() {
	dbCmd.AddCommand(dbMigrateCmd)
}
