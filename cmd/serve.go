/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/muharamdani/go-base-rest-api/migrations"
	"github.com/muharamdani/go-base-rest-api/seeders"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrating schema
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "This command will migrate all migrations func inside the migrations folder.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && args[0] == "rollback" {
			migrations.RollbackMigrations()
			fmt.Println("---Rollback success---")
			os.Exit(0)
		}
		migrations.ExecuteMigrations()
		fmt.Println("---Migration success---")
		os.Exit(0)
	},
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "This command will insert all seed data into the database. WARNING: Database data is deleted before seeding.",
	Run: func(cmd *cobra.Command, args []string) {
		seeders.Seed()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(seedCmd)
}
