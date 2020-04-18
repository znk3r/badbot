package db

import (
	log "github.com/sirupsen/logrus"
)

// GormLogger is a custom logger for Gorm using logrus
type GormLogger struct{}

// Print handles log events from Gorm sending them to logrus
func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		log.WithField("values", v[4]).Debug(v[3])
	case "log":
		log.WithField("source", v[1]).Warn(v[2])
	}
}
