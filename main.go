package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/paraizofelipe/gorecipes/handler"
	"github.com/paraizofelipe/gorecipes/settings"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	h := handler.New(logger)

	http.HandleFunc("/api/recipes/", h.RecipeHandler)

	url := fmt.Sprintf("%s:%s", settings.Host, settings.Port)

	log.Printf("🚀 Server listening in %s 🚀", url)

	if err := http.ListenAndServe(url, nil); err != nil {
		logger.Fatal(err)
	}
}
