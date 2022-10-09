package handlers

import (
	"context"
	"encoding/json"
	"github.com/altercolt/auth/pkg/web"
	"log"
	"net/http"
	"os"
)

func status(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	w.Header().Set("Content-Type", "application/json")
	status := map[string]string{
		"status": "ok",
	}

	res, err := json.Marshal(status)
	if err != nil {
		return err
	}

	w.Write(res)

	return nil
}

func Handlers(log *log.Logger, shutdown chan os.Signal) *web.App {
	app := web.NewApp(shutdown, nil)

	app.Handle(http.MethodGet, "/status", status)

	return app
}
