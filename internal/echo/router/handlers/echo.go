package handlers

import (
	"net/http"

	"github.com/erumble/go-echo/pkg/logger"
)

// EchoHandler is a handlerfunc that will echo back the reqquests made to it.
func EchoHandler(log logger.LeveledLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		echo(w, r, log)
	}
}

func echo(w http.ResponseWriter, r *http.Request, log logger.LeveledLogger) {
	log.Debugf("Echoing back request made to %s to client (%s)", r.URL.Path, r.RemoteAddr)

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Allow pre-flight headers
	w.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	if err := r.Write(w); err != nil {
		log.Error(err)
	}
}
