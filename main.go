package main

import (
	"awesomeProject/locale"
	log2 "awesomeProject/log"
	"awesomeProject/route"
	"awesomeProject/scheduler"
)

func main() {
	log2.Init()
	locale.Init()
	scheduler.Init()
	log := log2.InfoLogger
	log.Print("starting go app ...")
	app := route.App{}
	app.InitialiseRoutes()
	app.Run()
	log.Print("app started successfully")
}
