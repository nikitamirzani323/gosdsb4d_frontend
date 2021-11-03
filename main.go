package main

import (
	"log"

	"bitbucket.org/isbtotogroup/sdsb4d-frontend/routers"
)

func main() {
	app := routers.Init()
	log.Fatal(app.Listen(":80"))
}
