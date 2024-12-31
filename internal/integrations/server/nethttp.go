package server

import (
	"fmt"
	"net/http"

	"github.com/leagueify/account/handler"
	"github.com/leagueify/account/internal/config"
	"github.com/leagueify/account/internal/middleware"
)

type server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) Server {
	return &server{
		cfg: cfg,
	}
}

func (s *server) Start() {
	router := http.NewServeMux()
	accountRouter := handler.AccountRouter()

	router.Handle("/account/", http.StripPrefix("/account", accountRouter))

	// define server config
	accountServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Server.Port),
		Handler:      middleware.Logging(router),
		IdleTimeout:  s.cfg.Server.IdleTimeout,
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
	}

	// show server banner
	showStartBanner()

	// start server
	if err := accountServer.ListenAndServe(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
