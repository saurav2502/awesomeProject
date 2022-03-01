package main

import (
	log2 "awesomeProject/log"
	"awesomeProject/route"
)

func main() {
	log2.Init()
	log := log2.InfoLogger
	log.Print("starting go app ...")
	app := route.App{}
	app.InitialiseRoutes()
	app.Run()
	log.Print("app started successfully")
}
