/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/muharamdani/go-base-rest-api/migrations"
	"github.com/spf13/cobra"
	"os"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "migrate",
	Short: "This command will migrate all migrations func inside the migrations folder.",
	Run: func(cmd *cobra.Command, args []string) {
		migrations.ExecuteMigrations()
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
