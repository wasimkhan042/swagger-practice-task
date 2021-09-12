package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"custom-server/handlers"

	"custom-server/gen/restapi"
	"custom-server/gen/restapi/operations"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewSalamAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		_ = server.Shutdown()
	}()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// GetGreetingHandler greets the given name,
	// in case the name is not given, it will default to World
	api.GetGreetingHandler = handlers.Newgreeting()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}