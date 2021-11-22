package config

import (
	"fmt"
	"time"
)

const (
	defaultHost         = "0.0.0.0"
	defaultPort         = 8080
	defaultWaitShutdown = time.Second * 5
)

type (
	// Config - used as main config
	Config struct {
		Web Web
	}

	// Web - used as config for web server
	Web struct {
		Host         string
		Port         int
		WaitShutdown time.Duration
	}
)

// NewDefaultConfig - simple factory for create instance of Config
func NewDefaultConfig() *Config {
	return &Config{
		Web: Web{
			Host:         defaultHost,
			Port:         defaultPort,
			WaitShutdown: defaultWaitShutdown,
		},
	}
}

// Address - used for get web server address
func (w Web) Address() string {
	return fmt.Sprintf("%s:%d", w.Host, w.Port)
}
