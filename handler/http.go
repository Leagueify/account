package handler

import (
	"embed"
	"github.com/labstack/echo/v4"
	"github.com/leagueify/account/internal/config"
	"net/http"
)

var (
	//go:embed docs/*
	doc   embed.FS
	docFS = echo.MustSubFS(doc, "docs")
)

type httpHandler struct {
	app *echo.Echo
	cfg *config.Config
}

func (h *httpHandler) Initialize() {
	group := h.app.Group("/account")
	group.StaticFS("/docs", docFS)
	group.GET("/healthz", healthz)
}
func HTTP(app *echo.Echo, cfg *config.Config) Handler {
	return &httpHandler{
		app: app,
		cfg: cfg,
	}
}
func healthz(ctx echo.Context) error {
	type healthResponse struct {
		Service string `json:"service"`
		Status  string `json:"status"`
	}
	return ctx.JSON(http.StatusOK, healthResponse{
		Service: "account",
		Status:  http.StatusText(http.StatusOK),
	})
}
