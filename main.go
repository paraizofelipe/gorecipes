package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/paraizofelipe/gorecipes/handler"
	"github.com/paraizofelipe/gorecipes/settings"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(settings.Log)

	h := handler.New(e.Logger)
	v1 := e.Group("/api")
	h.Register(v1)

	fullHost := fmt.Sprintf("%s:%s", settings.Host, settings.Port)
	e.Logger.Fatal(e.Start(fullHost))
}
