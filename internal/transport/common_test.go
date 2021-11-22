package transport

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/yvv4git/go-echo/internal/config"
	"github.com/yvv4git/go-echo/internal/helpers"
	"github.com/yvv4git/go-echo/internal/infrastructure/logger"
	"github.com/yvv4git/go-echo/internal/services"
	"os"
	"testing"
)

type dependencies struct {
	logger     *zerolog.Logger
	cfg        *config.Config
	svcAddress *services.Address
	srvWeb     *Server
}

var testDeps dependencies

func loadDependencies() {
	log := logger.DefaultLogger()
	cfg := config.NewDefaultConfig()
	svcAddress := services.NewAddress(log, helpers.FilterIPv4Stub)
	srvWeb := NewServer(context.Background(), log, cfg, svcAddress)
	testDeps = dependencies{
		logger:     log,
		cfg:        cfg,
		svcAddress: svcAddress,
		srvWeb:     srvWeb,
	}
}

func TestMain(m *testing.M) {
	loadDependencies()
	code := m.Run()
	os.Exit(code)
}
