package cli

import (
	"linker/internal/database"
	"linker/internal/migration"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "drop application database",
	Run: func(cmd *cobra.Command, args []string) {
        db := database.GetDB()
        Drop(db)
	},
}

func init() {
	dbCmd.AddCommand(dropCmd)
}

func Drop(db *gorm.DB) {
	dbName := os.Getenv("POSTGRES_DB")

	println("Dropping database " + dbName + "...")

	for _, m := range migration.ModelList {
		if err := db.Migrator().DropTable(m); err != nil {
			log.Fatal(err)
		}
	}

	println("Dropping finished")
}
