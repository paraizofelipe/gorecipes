package handler

import (
	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/model"
)

type RecipeHandler interface {
	GetRecipes(echo.Context) error
	APIRecipeToRecipe([]model.APIRecipe) ([]model.Recipe, error)
}

// Handler represents a struct of handlers
type Handler struct {
	Logger echo.Logger
	Recipe RecipeHandler
}

// NewHandler ---
func New(logger echo.Logger) *Handler {
	return &Handler{
		Logger: logger,
		Recipe: NewRecipe(logger),
	}
}
