package main

import (
	"log"

	"github.com/nikitamirzani323/gosdsb4d_frontend/routers"
)

func main() {
	app := routers.Init()
	log.Fatal(app.Listen(":7071"))
}
