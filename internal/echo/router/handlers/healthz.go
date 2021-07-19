package handlers

import (
	"net/http"

	"github.com/erumble/go-echo/pkg/logger"
	"github.com/hellofresh/health-go/v4"
)

// HealthHandler is a handlerfunc that returns 200 if the service is healthy, 503 otherwise.
func HealthHandler(log logger.LeveledLogger) http.HandlerFunc {
	h, err := health.New()

	if err != nil {
		log.Errorf("Error instantiating health check: %s", err)
	}

	return h.HandlerFunc
}
