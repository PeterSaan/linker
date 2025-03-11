package cli

import "github.com/spf13/cobra"

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "APP_KEY commands for Linker",
}

func init() {
	rootCmd.AddCommand(keyCmd)
}
