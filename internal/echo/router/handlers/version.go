package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/erumble/go-echo/pkg/cli"
	"github.com/erumble/go-echo/pkg/logger"
)

type versionResp struct {
	Version string `json:"version"`
}

// VersionHandler is a handlerfunc that will return the version of the application.
func VersionHandler(log logger.LeveledLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		version(w, r, log)
	}
}

func version(w http.ResponseWriter, r *http.Request, log logger.LeveledLogger) {
	resp := versionResp{
		Version: cli.Version(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
