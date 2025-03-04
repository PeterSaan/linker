package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database commands for Linker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("db called")
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)

	dbCmd.PersistentFlags().StringP("test", "t", "testikas", "Testing flag")
}
