package handler

import "github.com/labstack/echo"

// Register ---
func (h Handler) Register(v1 *echo.Group) {
	recipes := v1.Group("/recipes")
	recipes.GET("", h.Recipes)
}
