package transport

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/yvv4git/go-echo/internal/config"
	"github.com/yvv4git/go-echo/internal/services"
	"net/http"
)

// Server - used as web server
type Server struct {
	logger     *zerolog.Logger
	router     *mux.Router
	web        *http.Server
	svcAddress *services.Address
}

// NewServer - simple factory for create instance of Server
func NewServer(ctx context.Context, logger *zerolog.Logger, cfg *config.Config, svcAddress *services.Address) *Server {
	server := &Server{
		logger:     logger,
		router:     mux.NewRouter(),
		svcAddress: svcAddress,
	}

	server.router.HandleFunc("/health", server.healthCheck).Methods(http.MethodGet)
	server.router.HandleFunc("/v1/address", server.addressHandlerV1).Methods(http.MethodGet)

	server.web = &http.Server{
		Addr:    cfg.Web.Address(),
		Handler: server.router,
	}

	go func() {
		<-ctx.Done()
		ctxShutdown, cancel := context.WithTimeout(context.Background(), cfg.Web.WaitShutdown)
		defer cancel()
		if err := server.Stop(ctxShutdown); err != nil {
			logger.Error().Err(err).Msg("error on shutdown web server")
		}
	}()

	return server
}

// Start - used for start web server
func (s Server) Start() error {
	s.logger.Info().Msg("start web server")
	return s.web.ListenAndServe()
}

// Stop - used for stop web server
func (s Server) Stop(ctx context.Context) error {
	s.logger.Info().Msg("stop web server")
	return s.web.Shutdown(ctx)
}
