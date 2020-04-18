package db

import (
	"github.com/jinzhu/gorm"
	// blank import is required to load the db driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	log "github.com/sirupsen/logrus"
)

// Connect opens the connection to the SQLite database
func Connect(file string) *gorm.DB {
	conn, err := gorm.Open("sqlite3", file)
	if err != nil {
		log.WithError(err).Fatalf("Failed to connect to the database %s", file)
	}
	log.Debugf("Connection to database %s established", file)

	conn.BlockGlobalUpdate(true)
	conn.SetLogger(&GormLogger{})
	conn.LogMode(true)

	return conn
}

// Disconnect closses the connection to the SQLite database
func Disconnect(conn *gorm.DB) {
	conn.Close()
	log.Debug("Connection to database closed")
}
