package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Tag model
type Tag struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `gorm:"unique_index;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Songs     []*Song `gorm:"many2many:song_tags;"`
}

// BeforeCreate allows to set or update values before the row is added
func (user *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New())
	return nil
}
