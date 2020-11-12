package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/paraizofelipe/gorecipes/handler"
)

var (
	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	h := handler.New(logger)

	http.HandleFunc("/api/recipes/", h.RecipeHandler)

	url := fmt.Sprintf("%s:%s", HOST, PORT)

	log.Printf("🚀 Server listening in %s 🚀", url)

	if err := http.ListenAndServe(url, nil); err != nil {
		logger.Fatal(err)
	}
}
