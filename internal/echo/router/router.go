package router

import (
	"net/http"

	"github.com/erumble/go-echo/internal/echo/router/handlers"
	"github.com/erumble/go-echo/pkg/logger"
	"github.com/gorilla/mux"
)

// Router allows us to pass in middleware
type Router interface {
	http.Handler
	WithMiddleware(middleware ...mux.MiddlewareFunc)
}

type router struct {
	*mux.Router
}

// New registers the routes and middleware for the server and returns an http handler
func New(log logger.LeveledLogger) Router {
	r := mux.NewRouter()
	r.HandleFunc("/status", handlers.HealthHandler(log))
	// r.HandleFunc("/static", handlers.StaticResponseHandler(staticResponse, log))
	r.HandleFunc("/version", handlers.VersionHandler(log))
	r.PathPrefix("/").Handler(handlers.EchoHandler(log))
	return &router{r}
}

func (r *router) WithMiddleware(middleware ...mux.MiddlewareFunc) {
	r.Use(middleware...)
}
