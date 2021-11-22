package transport

import (
	"encoding/json"
	"net/http"
)

const msgErrWriteBody = "can't write body"

func (s Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (s Server) addressHandlerV1(w http.ResponseWriter, r *http.Request) {
	result := s.svcAddress.IPv4()
	response, err := json.Marshal(result)
	if err != nil {
		s.logger.Error().Err(err).Msg("error on marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		s.logger.Err(err).Msg(msgErrWriteBody)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
