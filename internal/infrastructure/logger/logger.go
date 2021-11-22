package logger

import (
	"github.com/rs/zerolog"
	"os"
)

// DefaultLogger - used for create instance of default logger
func DefaultLogger() *zerolog.Logger {
	zerolog.TimestampFieldName = "date"
	zerolog.TimeFieldFormat = "2006.01.02 15:05:05"
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	return &logger
}
