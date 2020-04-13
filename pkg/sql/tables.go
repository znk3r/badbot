package sql

import (
	log "github.com/sirupsen/logrus"
	"database/sql"
	// blank import is required to load the db driver
	_ "github.com/mattn/go-sqlite3"
)

// Version of the database
const dbVersion = 1

func validateVersion(db *sql.DB) {
	if checkIfTableExist(db, "badbot_db_version") {
		var version int
		row := db.QueryRow("SELECT version FROM badbot_db_version ORDER BY rowid DESC;")
		if err := row.Scan(&version); err != nil {
			log.WithError(err).Fatal("Error validating the DB version")
		}

		if dbVersion > version {
			log.Fatalf("You have an older version of the database. Found version %d, expecting version %d", version, dbVersion)
		} else if dbVersion < version {
			log.Fatalf("You have a newer version of the database. Found version %d, expecting version %d", version, dbVersion)
		} else {
			log.Debugf("DB: Expected database version %d", dbVersion)
			log.Debugf("DB: Detected database version %d", version)
		}
	} else {
		initVersionTable(db)
	}
}

func initVersionTable(db *sql.DB) {
	sqlTable := "CREATE TABLE IF NOT EXISTS badbot_db_version(version INTEGER NOT NULL UNIQUE);"
	if _, err := db.Exec(sqlTable); err != nil {
		log.WithError(err).Fatal("Error creating the badbot_db_version table")
	}

	tx, err := db.Begin()
	if err != nil {
		log.WithError(err).Fatal("Error creating transaction")
	}
	stmt, err := tx.Prepare("INSERT INTO badbot_db_version(version) VALUES(?);")
	if err != nil {
		log.WithError(err).Fatal("Error creating transaction")
	}
	defer stmt.Close()
	if _, err = stmt.Exec(dbVersion); err != nil {
		log.WithError(err).Fatal("Error versioning the database")
	}
	tx.Commit()
}

func initSongsTable(db *sql.DB) {
	if !checkIfTableExist(db, "songs") {
		log.Debug("DB: Creating songs table")
		sqlTable := `
		CREATE TABLE IF NOT EXISTS songs(
			song_id INTEGER PRIMARY KEY,
			filename TEXT NOT NULL UNIQUE,
			hash TEXT NOT NULL,
			status TEXT NOT NULL,
			title TEXT NOT NULL,
			theme TEXT,
			artist TEXT,
			album TEXT,
			duration INTEGER
		);
		`

		if _, err := db.Exec(sqlTable); err != nil { 
			log.WithError(err).Fatal("Error creating the songs table")
		}

		log.Debug("DB: Creating index songs_filename")
		sqlTable = "CREATE INDEX IF NOT EXISTS song_filename ON songs (filename);"
		if _, err := db.Exec(sqlTable); err != nil { 
			log.WithError(err).Error("Error creating the song_filename index")
		}
	}
}

func initPlaylistsTable(db *sql.DB) {
	if !checkIfTableExist(db, "playlists") {
		log.Debug("DB: Creating playlists table")
		sqlTable := `
		CREATE TABLE IF NOT EXISTS playlists(
			playlist_id INTEGER PRIMARY KEY,
			title TEXT NOT NULL UNIQUE
		);
		`
		
		if _, err := db.Exec(sqlTable); err != nil { 
			log.WithError(err).Fatal("Error creating the playlists table")
		}
	}

	if !checkIfTableExist(db, "playlist_songs") {
		log.Debug("DB: Creating playlist_songs table")
		sqlTable := `
		CREATE TABLE IF NOT EXISTS playlist_songs(
			song_id INTEGER,
			playlist_id INTEGER,
			FOREIGN KEY(song_id) REFERENCES songs(song_id),
			FOREIGN KEY(playlist_id) REFERENCES playlists(playlist_id)
		);
		`

		if _, err := db.Exec(sqlTable); err != nil { 
			log.WithError(err).Error("Error creating the playlist_songs table")
		}
	}
}

func initTagsTable(db *sql.DB) {
	if !checkIfTableExist(db, "tags") {
		log.Debug("DB: Creating tags table")
		sqlTable := `
		CREATE TABLE IF NOT EXISTS tags(
			tag_id INTEGER PRIMARY KEY, 
			name TEXT NOT NULL UNIQUE
		);
		`

		if _, err := db.Exec(sqlTable); err != nil { 
			log.WithError(err).Fatal("Error creating the tags table")
		}

		log.Debug("DB: Creating index tag_name")
		sqlTable = "CREATE INDEX IF NOT EXISTS tag_name ON tags (name);"
		if _, err := db.Exec(sqlTable); err != nil { 
			log.WithError(err).Error("Error creating the tag_name index")
		}
	}

	if !checkIfTableExist(db, "song_tags") {
		log.Debug("DB: Creating song_tags table")
		sqlTable := `
		CREATE TABLE IF NOT EXISTS song_tags(
			song_id INTEGER,
			tag_id INTEGER,
			FOREIGN KEY(song_id) REFERENCES songs(song_id),
			FOREIGN KEY(tag_id) REFERENCES tags(tag_id)
		);
		`

		if _, err := db.Exec(sqlTable); err != nil { 
			log.WithError(err).Error("Error creating the song_tags table")
		}
	}
}
