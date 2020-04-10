package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "badbot",
	Short: "Discord bot to play music during online RPG sessions",
	Long: fmt.Sprintf(printBanner() + `BadBot is a discord bot, controlled with a SPA and an API. It
allows you to have a tagged music library to play background music
during online table RPG sessions.`),
	Version: Version,
	Run: serverHandler,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetVersionTemplate(fmt.Sprintf("BadBot version {{.Version}} (build %s from %s)\n", GitTag, BuildDate))

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "run in debug mode (slower)")
	rootCmd.PersistentFlags().StringP("port", "p", "8080", "port for the webserver")
	rootCmd.PersistentFlags().StringP("token", "t", "", "discord authentication token")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("server.port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("discord.auth_token", rootCmd.PersistentFlags().Lookup("token"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Print(printBanner())

	viper.AutomaticEnv() // read in environment variables that match

	if cfgFile == "" {
		cfgFile = selectDefaultConfigFile()
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
	}

	// Configure logging before continuing
	configureLogging()

	if cfgFile != "" {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}

func selectDefaultConfigFile() string {
	filenames := [...]string{"config.yaml", "config.yml", "config.json"}
	
	for _, file := range filenames {
		if _, err := os.Stat(file); !os.IsNotExist(err) {
			return file
		}
	}

	return ""
}

func printBanner() string {
	return fmt.Sprintf(`
	██████╗  █████╗ ██████╗ ██████╗  ██████╗ ████████╗
	██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔═══██╗╚══██╔══╝
	██████╔╝███████║██║  ██║██████╔╝██║   ██║   ██║   
	██╔══██╗██╔══██║██║  ██║██╔══██╗██║   ██║   ██║   
	██████╔╝██║  ██║██████╔╝██████╔╝╚██████╔╝   ██║   
	╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚═════╝  ╚═════╝    ╚═╝   
	`+"Version %s (build %s)\n\n", Version, GitTag)
}

func configureLogging() {
	debug := viper.GetBool("debug")
	if debug {
		log.SetLevel(log.DebugLevel)
	}
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: debug,
	})
}
