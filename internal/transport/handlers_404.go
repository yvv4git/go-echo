package transport

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"net/http"
	"strings"
)

// For trace 404 error
type noFoundHandler struct {
	logger *zerolog.Logger
}

// NewNoFoundHandler - used as simple factory for create noFoundHandler instance
func NewNoFoundHandler(logger *zerolog.Logger) *noFoundHandler {
	return &noFoundHandler{
		logger: logger,
	}
}

func (h *noFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseData := NewNoFoundResponse(r)
	response, err := json.Marshal(responseData)
	if err != nil {
		h.logger.Error().Err(err).Msg("error on marshaling response data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	_, err = w.Write(response)
	if err != nil {
		h.logger.Error().Err(err).Msg("error on write response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type noFoundResponse struct {
	Method string
	Host   string
	Path   string
}

// NewNoFoundResponse - used as simple factory for create noFoundResponse instance
func NewNoFoundResponse(r *http.Request) *noFoundResponse {

	return &noFoundResponse{
		Method: r.Method,
		Host:   r.Host,
		Path:   strings.TrimPrefix(r.URL.Path, "/provisions/"),
	}
}
