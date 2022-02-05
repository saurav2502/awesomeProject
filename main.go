package main

import (
	"awesomeProject/route"
	"log"
)

func main() {
	log.Print("starting go app ...")
	app := route.App{}
	app.InitialiseRoutes()
	app.Run()
	log.Print("app started successfully")
}
