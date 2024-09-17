package main

import (
	"github.com/Quorum-Code/portfolio-site-go/internal/webserver"
	"github.com/lpernett/godotenv"
)

func main() {
	godotenv.Load(".env")
	webserver.StartWebServer()
}
