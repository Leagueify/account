package main

import (
	"github.com/leagueify/account/internal/config"
	"github.com/leagueify/account/internal/integrations/server"
)

func main() {
	cfg := config.GetConfig()

	server.NewServer(cfg).Start()
}
