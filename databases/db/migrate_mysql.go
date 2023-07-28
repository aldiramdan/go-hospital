package db

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "db migration",
	RunE:  dbMigrate,
}

var migUp, migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", true, "run migration up")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db := ConnectDB()
	if migDown {
		log.Println("Migration down done")
		return executeMigrationDown(db)
	}

	if migUp {
		log.Println("Migration up done")
		return executeMigrationUp(db)
	}

	return nil
}

func executeMigrationDown(db *sql.DB) error {
	file, err := os.ReadFile(os.Getenv("PathMigrateDown"))
	if err != nil {
		return err
	}

	err = excuteMigrations(db, file)
	if err != nil {
		return err
	}

	return nil
}

func executeMigrationUp(db *sql.DB) error {
	file, err := os.ReadFile(os.Getenv("PathMigrateUp"))
	if err != nil {
		return err
	}

	err = excuteMigrations(db, file)
	if err != nil {
		return err
	}

	return nil
}

func excuteMigrations(db *sql.DB, file []byte) error {
	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			return err
		}
	}

	return nil
}
