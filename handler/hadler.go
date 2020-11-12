package handler

import "log"

// ErrorResponse ---
type ErrorResponse struct {
	Status int                    `json:"-"`
	Errors map[string]interface{} `json:"errors"`
}

// Handler ---
type Handler struct {
	Logger *log.Logger
}

// NewHandler ---
func New(logger *log.Logger) *Handler {
	return &Handler{
		Logger: logger,
	}
}
