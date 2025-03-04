package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dbMigrateCmd = &cobra.Command{
	Use:   "db migrate",
	Short: "Generate database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("db migrate called")
	},
}

func init() {
	dbCmd.AddCommand(dbMigrateCmd)

}
