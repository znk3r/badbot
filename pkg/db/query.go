package db

import (
	log "github.com/sirupsen/logrus"
	"github.com/znk3r/badbot/pkg/model"
)

// GetSongs returns an array with all the songs matching the filter
func (c *DbConn) GetSongs() []model.Song {
	var songs []model.Song
	if err := c.conn.Find(&songs).Error; err != nil {
		log.WithError(err).Error("Error reading from the songs table")
	}

	return songs
}

// CountSongs returns the number of songs matching the filter
func (c *DbConn) CountSongs() int {
	var count int
	c.conn.Model(&model.Song{}).Count(&count)

	return count
}

// CountPlaylists returns the number of playlists matching the filter
func (c *DbConn) CountPlaylists() int {
	var count int
	c.conn.Model(&model.Playlist{}).Count(&count)

	return count
}

// CountTags returns the number of tags matching the filter
func (c *DbConn) CountTags() int {
	var count int
	c.conn.Model(&model.Tag{}).Count(&count)

	return count
}
