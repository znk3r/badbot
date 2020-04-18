package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/znk3r/badbot/pkg/db"
)

// dbInitCmd represents the db:init command
var dbInitCmd = &cobra.Command{
	Use:        "db:init",
	Aliases:    []string{"db:new", "db:migrate"},
	SuggestFor: []string{"db"},
	Short:      "Create or update the database",
	Long: printBanner() +
		`Create and initialize a new database for badbot, or migrate an existing one.

If the database file doesn't exist, this command will create a new one with the
expected structure. If a database file already exists, it'll try to migrate the
tables.

A database migration will only add missing tables, columns and indexes, but won't
change existing column types or delete unused columns to protect the existing
data.`,
	Example: `  You can create the default DB with:
    badbot db:init
  Or use a specific file like "my-storage.db" with:
    badbot db:init -d my-storage.db`,
	Run: dbInitHandler,
	PreRun: func(dbInitCmd *cobra.Command, args []string) {
		// Workaround for known bug in viper+cobra
		// Viper is declared globally, so you can't have the same parameter
		// name for two different commands
		// https://github.com/spf13/viper/issues/233
		viper.BindPFlag("db.file", dbInitCmd.Flags().Lookup("db"))
	},
}

func init() {
	rootCmd.AddCommand(dbInitCmd)

	dbInitCmd.Flags().StringP("db", "d", "data.db", "SQLite file to use as database")
}

func dbInitHandler(cmd *cobra.Command, args []string) {
	file := viper.GetString("db.file")
	log.Infof("Using database: %s", file)

	conn := db.Connect(file)
	defer conn.Disconnect()

	conn.MigrateDatabase()
	log.Infof("Database %s updated", file)
}
