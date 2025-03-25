package cli

import (
	"linker/internal/database"
	"linker/internal/migration"
	"linker/internal/seeder"
	"log"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var Seed bool
var Fresh bool

var dbMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Generate database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		println("Connecting to database")
		db := database.GetDB()

		if Fresh {
			migrateFresh(db)

			return
		}

        migrate(db)

		if Seed {
			migrateSeed()
		}
	},
}

func init() {
	dbMigrateCmd.Flags().BoolVarP(&Seed, "seed", "s", false, "seed the database")
	dbMigrateCmd.Flags().BoolVarP(&Fresh, "fresh", "f", false, "drop database and seed")

	dbCmd.AddCommand(dbMigrateCmd)
}

func migrate(db *gorm.DB) {
	println("Migrating...")

	if err := migration.Migrate(db); err != nil {
		log.Fatal(err)
	}

	println("Migration finished")
}

func migrateSeed() {
	println("Seeding...")

	if err := seeder.Main(); err != nil {
		log.Fatal(err)
	}

	println("Seeding finished")
}

func migrateFresh(db *gorm.DB) {
    Drop(db)

    migrate(db) 

	println("Migration finished")

	println("Seeding...")

	if err := seeder.Main(); err != nil {
		log.Fatal(err)
	}

	println("Seeding finished")
}
