package handler

import (
	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/external"
)

// Handler represents a struct of handlers
type Handler struct {
	Logger      echo.Logger
	ExternalAPI external.API
}

// NewHandler ---
func New(logger echo.Logger) *Handler {
	return &Handler{
		Logger:      logger,
		ExternalAPI: external.New(logger),
	}
}
