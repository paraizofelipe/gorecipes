package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/handler"
	"github.com/paraizofelipe/gorecipes/settings"
)

func main() {
	e := echo.New()

	h := handler.New(e.Logger)
	v1 := e.Group("/api")
	h.Register(v1)

	e.GET("/api/recipes", h.Recipes)
	fullHost := fmt.Sprintf("%s:%s", settings.Host, settings.Port)
	// log.Printf("🚀 Server listening in %s 🚀", url)
	e.Logger.Fatal(e.Start(fullHost))
}
