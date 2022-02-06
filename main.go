package main

import (
	"awesomeProject/route"
	"log"
)

func main() {
	log.Print("starting go app ...")
	// Set the file name of the configurations file
	/*viper.SetConfigName("config")

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
	}*/

	// Reading variables using the model
	/*fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configuration.Database.DBName)
	fmt.Println("Port is\t\t", configuration.Server.Port)*/
	//fmt.Println("EXAMPLE_PATH is\t", configuration.EXAMPLE_PATH)
	//fmt.Println("EXAMPLE_VAR is\t", configuration.EXAMPLE_VAR)
	app := route.App{}
	app.InitialiseRoutes()
	app.Run()
	log.Print("app started successfully")
}
