package cli

import "github.com/spf13/cobra"

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database commands for Linker",
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
