package discord

import (
	"os"
	
	log "github.com/sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
)

// StartBot starts the discord session and the bot
func StartBot(token string) *discordgo.Session {
	if (token == "") {
		log.Error("You must provide a Discord authentication token")
		os.Exit(1)
	}

	// Start discord session
	log.Debug("Creating discord session")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.WithError(err).Error("Error creating Discord session")
		os.Exit(1)
	}
	log.Debug("Bot authorized, session created")

	// Open a websocket connection to Discord
	log.Debug("Opening connection to discord")
	err = session.Open()
	if err != nil {
		log.WithError(err).Error("Error opening connection to Discord")
		os.Exit(1)
	}
	log.Debug("Connected to discord")

	return session
}

// KillBot closes the discord session
func KillBot(session *discordgo.Session) {
	// Clean up
	log.Debug("Clossing connection to discord")
	session.Close()
	log.Debug("Disconnected from discord")
}
