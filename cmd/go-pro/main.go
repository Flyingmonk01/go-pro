package main

import (
	// formatting and printing

	"log/slog"
	"net/http"
	"os"

	"github.com/Flyingmonk01/go-pro/internal/router"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	slog.Info("🚀 Server starting", "port", ":8080")

	r := router.SetupRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error("server crashed", "error", err)
	}
}
