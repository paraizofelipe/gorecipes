package settings

import "os"

var (
	Host      string
	Port      string
	GyphToken string
)

func init() {
	Host = os.Getenv("HOST")
	Port = os.Getenv("PORT")
	GyphToken = os.Getenv("GYPH_TOKEN")
}
