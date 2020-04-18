package db

import (
	"github.com/jinzhu/gorm"
	"github.com/znk3r/badbot/pkg/model"
	// blank import is required to load the db driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	log "github.com/sirupsen/logrus"
)

// DbConn is the structure holding the database connection
type DbConn struct {
	conn        *gorm.DB
	isConnected bool
}

// Connect opens the connection to the SQLite database
func Connect(file string) *DbConn {
	conn, err := gorm.Open("sqlite3", file)
	if err != nil {
		log.WithError(err).Fatalf("Failed to connect to the database %s", file)
	}
	log.Debugf("Connection to database %s established", file)

	conn.BlockGlobalUpdate(true)
	conn.SetLogger(&GormLogger{})
	conn.LogMode(true)

	return &DbConn{
		conn:        conn,
		isConnected: true,
	}
}

// Disconnect closses the connection to the SQLite database
func (c *DbConn) Disconnect() {
	if !c.isConnected {
		log.Warn("Database already disconnected")
	}

	c.isConnected = false
	c.conn.Close()
	log.Debug("Connection to database closed")
}

func (c *DbConn) MigrateDatabase() {
	c.conn.AutoMigrate(
		&model.Song{},
		&model.Playlist{},
		&model.Tag{},
	)
}
