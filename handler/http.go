package handler

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	//go:embed docs/*
	doc   embed.FS
	docFS = echo.MustSubFS(doc, "docs")
)

type httpHandler struct {
	app *echo.Echo
}

func (h *httpHandler) Initialize() {
	group := h.app.Group("/account")
	group.StaticFS("/docs", docFS)
	group.GET("/healthz", healthz)
}
func HTTP(app *echo.Echo) Handler {
	return &httpHandler{
		app: app,
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
