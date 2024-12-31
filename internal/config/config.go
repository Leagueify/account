package config

import (
	"fmt"
	"os"
	"time"
)

type (
	Config struct {
		Server
	}

	Server struct {
		Port         int
		IdleTimeout  time.Duration
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
)

func GetConfig() *Config {
	config := &Config{}
	config.getDefaults()
	config.loadFromEnv()
	return config
}

func (cfg *Config) getDefaults() {
	// server defaults
	cfg.Server.Port = 6501
	cfg.Server.IdleTimeout = time.Minute
	cfg.Server.ReadTimeout = 10 * time.Second
	cfg.Server.WriteTimeout = 30 * time.Second
}

func (cfg *Config) loadFromEnv() {
	// server config
	if cfgVar := os.Getenv("SERVER_IDLE_TIMEOUT"); cfgVar != "" {
		idleTimeout, err := time.ParseDuration(cfgVar)
		if err != nil {
			panic(
				fmt.Sprintf(
					"CONFIG IDLE TIMEOUT ERROR: %v\n",
					err,
				),
			)
		}
		cfg.Server.IdleTimeout = idleTimeout
	}
	if cfgVar := os.Getenv("SERVER_READ_TIMEOUT"); cfgVar != "" {
		readTimeout, err := time.ParseDuration(cfgVar)
		if err != nil {
			panic(
				fmt.Sprintf(
					"CONFIG READ TIMEOUT ERROR: %v\n",
					err,
				),
			)
		}
		cfg.Server.IdleTimeout = readTimeout
	}
	if cfgVar := os.Getenv("SERVER_WRITE_TIMEOUT"); cfgVar != "" {
		writeTimeout, err := time.ParseDuration(cfgVar)
		if err != nil {
			panic(
				fmt.Sprintf(
					"CONFIG READ TIMEOUT ERROR: %v\n",
					err,
				),
			)
		}
		cfg.Server.IdleTimeout = writeTimeout
	}
}
