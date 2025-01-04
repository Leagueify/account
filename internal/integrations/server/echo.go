package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leagueify/account/handler"
	"github.com/leagueify/account/internal/config"
	"time"
)

type echoServer struct {
	app *echo.Echo
	cfg *config.Config
}

func NewEchoServer(cfg *config.Config) Server {
	echoApp := echo.New()
	echoApp.HideBanner = true
	return &echoServer{
		app: echoApp,
		cfg: cfg,
	}
}
func (s *echoServer) Start() {
	// configure middleware
	s.app.Use(middleware.RequestLoggerWithConfig(
		middleware.RequestLoggerConfig{
			LogLatency: true,
			LogHost:    true,
			LogMethod:  true,
			LogStatus:  true,
			LogURI:     true,
			LogValuesFunc: func(
				ctx echo.Context,
				v middleware.RequestLoggerValues,
			) error {
				fmt.Printf(
					"%v -- %v %v %v%v (%d ms)\n",
					time.Now().Format(time.Stamp), v.Status,
					v.Method, v.Host, v.URI,
					v.Latency.Abs().Milliseconds(),
				)
				return nil
			},
		},
	))
	s.app.Use(middleware.Recover())
	// register http handler
	handler.HTTP(s.app).Initialize()
	showStartBanner()
	// start server
	s.app.Logger.Fatal(s.app.Start(fmt.Sprintf(":%d", s.cfg.Server.Port)))
}
