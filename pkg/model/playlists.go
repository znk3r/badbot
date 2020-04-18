package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Playlist model
type Playlist struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title     string    `gorm:"unique_index;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Songs     []*Song `gorm:"many2many:song_playlists;"`
}

// BeforeCreate allows to set or update values before the row is added
func (user *Playlist) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
