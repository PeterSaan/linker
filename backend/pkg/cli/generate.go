package cli

import (
	"linker/internal/key"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate and set APP_KEY for Linker",
	Run: func(cmd *cobra.Command, args []string) {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}

		println("Generating key...")
		appKey, err := key.Generate()
		if err != nil {
			log.Fatal(err)
		}

		println("Setting APP_KEY to .env")
		envVars, err := godotenv.Read(".env")
		if err != nil {
			println(err)
		}
		envVars["LINKEDIN_OAUTH_CLIENT_SECRET"] = ""
		envVars["APP_KEY"] = appKey
		if err := godotenv.Write(
			envVars,
			".env",
		); err != nil {
			log.Fatal(err)
		}
		
		println("APP_KEY generated and set")
	},
}

func init() {
	keyCmd.AddCommand(generateCmd)
}
