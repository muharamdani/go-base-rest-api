package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	conn "github.com/muharamdani/go-base-rest-api/db"
	"github.com/muharamdani/go-base-rest-api/db/migrations"
	"github.com/muharamdani/go-base-rest-api/utils"

	"github.com/spf13/cobra"
)

var migrationTemplate = `package migrations

import (
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type TableName struct {}

func init() {
	// Fill this with your migrations
	migrations = []*gormigrate.Migration{}
}

// Migrate will execute users migrations
func Migrate(db *gorm.DB) {
	m = gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	fmt.Println("Migrated table")
}

// Rollback will rollback users migrations
func Rollback(db *gorm.DB) {
	m = gormigrate.New(db, gormigrate.DefaultOptions, migrations)
	if err := m.RollbackLast(); err != nil {
		log.Fatalf("Could not rollback: %v", err)
	}
	fmt.Println("Rolled back table")
}
`

// migrateCmd represents the migrating schema
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "This command will migrate all migrations func inside the migrations folder.",
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("create")
		if fileName != "" {
			if err := createMigrationFile(fileName); err != nil {
				log.Fatalf("Error creating migration file: %v", err)
				return
			}
			fmt.Printf("Migration file %s created!\n", fileName)
			os.Exit(0)
		}
		conn.Connect("default")
		if len(args) > 0 && args[0] == "rollback" {
			migrations.RollbackMigrations(conn.DB)
			fmt.Println("---Rollback success---")
			os.Exit(0)
		}
		migrations.ExecuteMigrations(conn.DB)
		fmt.Println("---Migration success---")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.PersistentFlags().String("create", "", "Create migration file")
}

func createMigrationFile(filename string) error {
	rootPath := utils.GetRootPath()
	t := time.Now()
	currTime := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	fileName := fmt.Sprintf("%s/db/migrations/%s_%s.go", rootPath, currTime, filename)
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err2 := f.WriteString(migrationTemplate)

	if err2 != nil {
		return err
	}

	return nil
}
