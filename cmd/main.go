package main

import (
	"context"
	"github.com/yvv4git/go-echo/internal/config"
	"github.com/yvv4git/go-echo/internal/helpers"
	"github.com/yvv4git/go-echo/internal/infrastructure/logger"
	"github.com/yvv4git/go-echo/internal/services"
	"github.com/yvv4git/go-echo/internal/transport"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := logger.DefaultLogger()
	cfg := config.NewDefaultConfig()
	svcAddress := services.NewAddress(log, helpers.FilterIPv4)
	srvWeb := transport.NewServer(ctx, log, cfg, svcAddress)

	go func() {
		termCh := make(chan os.Signal, 1)
		signal.Notify(termCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-termCh
		log.Info().Msg("got terminate signal")
		cancel()
	}()

	if err := srvWeb.Start(); err != nil && err != http.ErrServerClosed {
		log.Error().Err(err).Msg("error on start web server")
	}

	log.Info().Msg("server is successfully turned off")
}
