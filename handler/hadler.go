package handler

import (
	"github.com/labstack/echo"
)

// ErrorResponse ---
type ErrorResponse struct {
	Status int                    `json:"-"`
	Errors map[string]interface{} `json:"errors"`
}

// Handler ---
type Handler struct {
	Logger echo.Logger
}

// NewHandler ---
func New(logger echo.Logger) *Handler {
	return &Handler{
		Logger: logger,
	}
}
