package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Song model
type Song struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	FileName  string    `gorm:"unique_index;not null"`
	Hash      string    `gorm:"unique_index;not null"`
	Status    string    `gorm:"not null"`
	Theme     string
	Title     string `gorm:"not null"`
	Artist    string
	Album     string
	Duration  sql.NullInt64
	CreatedAt time.Time
	UpdatedAt time.Time
	Playlists []*Playlist `gorm:"many2many:song_playlists;"`
	Tags      []*Tag      `gorm:"many2many:song_tags;"`
}

// BeforeCreate allows to set or update values before the row is added
func (user *Song) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
