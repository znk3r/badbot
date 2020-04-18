package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/znk3r/badbot/pkg/db"
)

// dbInfoCmd represents the db:info command
var dbInfoCmd = &cobra.Command{
	Use:     "db:info",
	Aliases: []string{"db:stats"},
	Short:   "Show statistics from the database",
	Long: printBanner() +
		`Show statistics and information about the existing database.`,
	Run: dbInfoHandler,
	PreRun: func(dbInfoCmd *cobra.Command, args []string) {
		// Workaround for known bug in viper+cobra
		// Viper is declared globally, so you can't have the same parameter
		// name for two different commands
		// https://github.com/spf13/viper/issues/233
		viper.BindPFlag("db.file", dbInfoCmd.Flags().Lookup("db"))
	},
}

func init() {
	rootCmd.AddCommand(dbInfoCmd)

	dbInfoCmd.Flags().StringP("db", "d", "data.db", "SQLite file to use as database")
}

func dbInfoHandler(cmd *cobra.Command, args []string) {
	file := viper.GetString("db.file")
	log.Infof("Using database: %s", file)

	conn := db.Connect(file)
	defer conn.Disconnect()

	log.Info("Number of elements stored in the database:")
	log.Infof("  - %d songs", conn.CountSongs())
	log.Infof("  - %d playlists", conn.CountPlaylists())
	log.Infof("  - %d tags", conn.CountTags())
}
