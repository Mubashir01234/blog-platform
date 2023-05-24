package main

import (
	"log"
	"net/http"

	"blog/config" // Importing the config package
	"blog/routes" // Importing the routes package

	"github.com/fatih/color" // Importing the color package from the fatih library
)

// This is init function will run before main function and load the configs
func init() {
	err := config.LoadConfig() // Loading the configuration
	if err != nil {
		log.Fatal(err) // Exiting the program if there is an error loading the configuration
	}
}

// Main function will start the application
func main() {
	color.Cyan("🌏 Server running on localhost:" + config.Cfg.ServerPort) // Printing a colored message indicating the server is running

	// Starting the HTTP server on the specified port, using the routes defined in the routes package
	log.Fatal(http.ListenAndServe(":"+config.Cfg.ServerPort, routes.NewRoutesImpl()))
}
