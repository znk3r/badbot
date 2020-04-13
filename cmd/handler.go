package cmd

import (
	"context"
	goSql "database/sql"
	"fmt"
	"math"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/znk3r/badbot/pkg/api"
	"github.com/znk3r/badbot/pkg/sql"
	// "github.com/znk3r/badbot/pkg/discord"
)

func serverHandler(cmd *cobra.Command, args []string) {
	startDB()
	srv := startServer()

	// Handle Shutdown
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)
	// Block until we receive our signal.
	<-c
	fmt.Print("\n")

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), serverTimeout())
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	log.Debug("Shutting down web server")
	srv.Shutdown(ctx)

	log.Info("Shutting down")
}

func startDB() *goSql.DB {
	dbfile := viper.GetString("db.file")

	// Open connection to the database
	db := sql.Connect(dbfile)
	defer db.Close()

	// Check and initialize database
	sql.InitializeTables(db)

	return db
}

func startServer() *http.Server {
	port := viper.GetString("server.port")
	writeTimeout := viper.GetInt("server.write_timeout")
	readTimeout := viper.GetInt("server.read_timeout")
	idleTimeout := viper.GetInt("server.idle_timeout")

	router := api.CreateRoutes()

	srv := &http.Server{
		Addr: ":" + port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * time.Duration(writeTimeout),
		ReadTimeout:  time.Second * time.Duration(readTimeout),
		IdleTimeout:  time.Second * time.Duration(idleTimeout),
		Handler:      router,
	}

	fmt.Printf(`
API server running. You can now connect with your browser to

	Local: http://localhost:%s

Now running. Press CTRL-C to exit.`+"\n\n", port)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.WithError(err).Fatal("Error running the web server")
		}
	}()

	return srv
}

func serverTimeout() time.Duration {
	writeTimeout := viper.GetInt("server.write_timeout")
	readTimeout := viper.GetInt("server.read_timeout")
	timeout := math.Max(float64(writeTimeout), float64(readTimeout))

	return time.Duration(timeout)
}
