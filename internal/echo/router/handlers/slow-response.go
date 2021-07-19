package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/erumble/go-echo/pkg/logger"
)

type slowResp struct {
	Response string `json:"response"`
}

// StaticResponseHandler is a handlerfunc that will always return the same response.
func SlowResponseHandler(delay int, log logger.LeveledLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slowResponse(w, r, delay, log)
	}
}

func slowResponse(w http.ResponseWriter, r *http.Request, delay int, log logger.LeveledLogger) {
	log.Debug("Slow response called")

	startTime := time.Now()
	time.Sleep(time.Duration(delay) * time.Second)

	resp := slowResp{
		Response: fmt.Sprintf("response took %v seconds", time.Now().Sub(startTime).Seconds()),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
