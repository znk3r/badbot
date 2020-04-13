package sql

import (
	"fmt"
	
	log "github.com/sirupsen/logrus"
	"database/sql"
	// blank import is required to load the db driver
	_ "github.com/mattn/go-sqlite3"
)

// Connect to the SQLite database
func Connect(file string) *sql.DB {
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", file)

	log.Infof("Connecting to %s", file)
	log.Debugf("DB: DSN %s", dsn)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.WithError(err).Fatal("Error opening the database")
	}
	if db == nil {
		log.Fatal("Error opening the database, db is nil")
	}

	log.Debugf("DB: Connected successfully")
	return db
}

// InitializeTables creates the tables if they don't exist
func InitializeTables(db *sql.DB) {
	validateVersion(db)

	initSongsTable(db)
	initPlaylistsTable(db)
	initTagsTable(db)
}

func checkIfTableExist(db *sql.DB, table string) bool {
	var count int

	row := db.QueryRow("select count(name) from sqlite_master where type='table' AND name=$1;", table)
	err := row.Scan(&count)

	if err == sql.ErrNoRows || count == 0 {
		return false
	} else if err != nil {
		log.WithError(err).Error("Error checking if a table exists")
		return false
	}

	return true
}
