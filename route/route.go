package route

import (
	"awesomeProject/config"
	"awesomeProject/service"
	_ "awesomeProject/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
}

func (app *App) Run() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Port is\t\t", configuration.Server.Port)
	fmt.Println("DBURL is\t\t", configuration.Database.DbUrl)
	port := ":" + strconv.Itoa(configuration.Server.Port)
	log.Print("App is running on the port ", port)
	log.Fatal(http.ListenAndServe(port, app.Router))
}

func (app *App) InitialiseRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/", service.Handle)
}
