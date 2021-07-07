package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/erumble/go-echo/pkg/logger"
)

type staticResp struct {
	Response string `json:"response"`
}

// StaticResponseHandler is a handlerfunc that will always return the same response.
func StaticResponseHandler(response string, log logger.LeveledLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		staticResponse(w, r, response, log)
	}
}

func staticResponse(w http.ResponseWriter, r *http.Request, response string, log logger.LeveledLogger) {
	resp := staticResp{
		Response: response,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
