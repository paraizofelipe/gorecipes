package settings

import (
	"os"

	"github.com/labstack/gommon/log"
)

var (
	Host       string
	Port       string
	GiphyToken string
	Log        log.Lvl
)

func init() {
	logLvl := map[string]log.Lvl{
		"ERROR": log.ERROR,
		"DEBUG": log.DEBUG,
		"INFO":  log.INFO,
	}

	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	GiphyToken = os.Getenv("GIPHY_TOKEN")
	Log = logLvl[os.Getenv("LOG")]
}
