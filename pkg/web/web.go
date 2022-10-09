package web

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"syscall"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	mux         *mux.Router
	shutdown    chan os.Signal
	middlewares []Middleware
}

func NewApp(shutdown chan os.Signal, middlewares ...Middleware) *App {
	app := App{
		mux:         mux.NewRouter(),
		shutdown:    shutdown,
		middlewares: middlewares,
	}

	return &app
}

func (a *App) signalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

func (a *App) Handle(method string, path string, handler Handler, middlewares ...Middleware) {
	handler = wrapMiddleware(middlewares, handler)
	handler = wrapMiddleware(a.middlewares, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if err := handler(ctx, w, r); err != nil {
			a.signalShutdown()
			return
		}
	}

	a.mux.HandleFunc(path, h).Methods(method)
}
