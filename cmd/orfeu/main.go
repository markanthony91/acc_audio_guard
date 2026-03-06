package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/marcelo/acc_audio_guard/internal/app"
	"github.com/marcelo/acc_audio_guard/internal/config"
)

func main() {
	cfg := config.FromFlags()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("[Orfeu] bootstrap_error: %v", err)
	}

	if err := application.Run(ctx); err != nil {
		log.Fatalf("[Orfeu] runtime_error: %v", err)
	}
}
