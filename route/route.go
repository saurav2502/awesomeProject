package route

import (
	"awesomeProject/service"
	_ "awesomeProject/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":8083", app.Router))
}

func (app *App) InitialiseRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/", service.Handle)
}
