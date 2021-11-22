package services

import (
	"github.com/rs/zerolog"
	"github.com/yvv4git/go-echo/internal/helpers"
	"github.com/yvv4git/go-echo/internal/responses"
)

// Address - used for get IPv4 address of this service
type Address struct {
	logger   *zerolog.Logger
	filterIP helpers.FilterIP
}

// IPv4 - return IPv4 address
func (a Address) IPv4() responses.Address {
	address, err := helpers.FindIP(a.filterIP)
	if err != nil {
		a.logger.Error().Err(err).Msg("error on get ip address")
	}
	return responses.Address{
		HostAddress: address,
	}
}

// NewAddress - simple factory for create instance of Address
func NewAddress(logger *zerolog.Logger, filterIP helpers.FilterIP) *Address {
	return &Address{
		filterIP: filterIP,
		logger:   logger,
	}
}
