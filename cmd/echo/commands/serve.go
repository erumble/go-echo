package commands

import (
	"context"
	golog "log"
	"strings"
	"time"

	"github.com/erumble/go-echo/internal/echo/router"
	"github.com/erumble/go-echo/pkg/cli"
	"github.com/erumble/go-echo/pkg/logger"
	"github.com/erumble/go-echo/pkg/middleware/httplogger"
	"github.com/erumble/go-echo/pkg/server"
)

type serveCmd struct {
	Port       int    `short:"p" long:"port" env:"PORT" default:"8080" required:"false" description:"The port on which the service listens"`
	StaticResp string `short:"r" long:"static-response" env:"STATIC_RESPONSE" required:"true" description:"Return value for the static response handler"`
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

	r := router.New(cmd.StaticResp, log)
	r.WithMiddleware(httplogger.HTTPLogger)
	s := server.New(r, cmd.Port, 5*time.Second, log)

	return s.Serve(ctx)
}
