package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	// log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/znk3r/badbot/pkg/discord"
)


func serverHandler(cmd *cobra.Command, args []string) {
	// port := viper.GetString("server.port")
	token := viper.GetString("discord.auth_token")

	session := discord.StartBot(token)

	// Wait for a CTRL-C
	fmt.Print("\nNow running. Press CTRL-C to exit.\n\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	fmt.Print("\n")

	// Clean up
	discord.KillBot(session)
}

