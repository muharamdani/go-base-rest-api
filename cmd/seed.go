package cmd

import (
	conn "github.com/muharamdani/go-base-rest-api/db"
	"github.com/muharamdani/go-base-rest-api/db/seeders"
	"os"

	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "This command will insert all seed data into the database. WARNING: Database data is deleted before seeding.",
	Run: func(cmd *cobra.Command, args []string) {
		conn.Connect("default")
		seeders.Seed(conn.DB)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
