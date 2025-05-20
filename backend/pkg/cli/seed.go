package cli

import (
	"linker/internal/seeder"
	"log"

	"github.com/spf13/cobra"
)

var dbSeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Run database seeder",
	Run: func(cmd *cobra.Command, args []string) {
		println("Seeding...")
		if err := seeder.Main(); err != nil {
			log.Fatal(err)
		}

		println("Seeder finished")
	},
}

func init() {
	dbCmd.AddCommand(dbSeedCmd)
}
