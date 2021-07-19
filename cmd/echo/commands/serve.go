package commands

import (
	"context"
	golog "log"
	"strings"
	"time"

	"github.com/erumble/go-echo/internal/echo/router"
	"github.com/erumble/go-echo/pkg/cli"
	"github.com/erumble/go-echo/pkg/logger"
	"github.com/erumble/go-echo/pkg/middleware"
	"github.com/erumble/go-echo/pkg/server"
)

type serveCmd struct {
	Port          int `short:"p" long:"port" env:"PORT" default:"8080" required:"false" description:"The port on which the service listens"`
	SlowRespDelay int `short:"d" long:"slow-response-delay" env:"SLOW_RESPONSE_DELAY" default:"5" required:"false" description:"Time in seconds that the slow response endpoint will wait before sending response"`
}

func init() {
	var cmd serveCmd
	if _, err := cli.AddCommand(
		&cmd,
		"serve",
		"Start server",
		"Start an HTTP echo server",
	); err != nil {
		// yep, panic. If this fails something is wrong with either the serveCmd struct, or the serveCmd.Execute function
		panic(err)
	}
}

// Execute implements the Commander interace from the jessevdk/go-flags package
// We don't (currnetly) care about positional arguments, so we use an `_ []string` to ignore them.
func (cmd serveCmd) Execute(_ []string) error {
	// Set up the logger
	log := logger.NewLeveledLogger(cli.GlobalOpts.LogLevel)
	defer func() {
		if err := log.Sync(); err != nil {
			// Derpy workaround for https://github.com/uber-go/zap/issues/880
			if !(strings.Contains(err.Error(), "/dev/stderr") && strings.Contains(err.Error(), "/dev/stdout")) {
				// Using golog since the error was with the logging package. Maybe it's okay to use the log.Error here, I dunno
				golog.Printf("Error syncing logs: %v\n", err)
			}
		}
	}()

	log.Debug("DEBUG logging enabled")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := router.New(cmd.SlowRespDelay, log)
	r.WithMiddleware(middleware.HTTPLogger)
	s := server.New(r, cmd.Port, 5*time.Second, log)

	return s.Serve(ctx)
}
