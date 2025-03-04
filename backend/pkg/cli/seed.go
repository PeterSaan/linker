package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dbSeedCmd = &cobra.Command{
	Use:   "db seed",
	Short: "Run database seeder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("db seed called")
	},
}

func init() {
	dbCmd.AddCommand(dbSeedCmd)
}
