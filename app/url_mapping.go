package app

import (
	"backendReq/logger"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func mapUrls() {

	//Assign Repository//

	brh := BackendReqHandler{}

	router.HandleFunc("/retrivelistofproduct/{customer_id:[0-9]+}", brh.GetRetriveListOfProduct).Methods(http.MethodGet).
		Name("GetRetriveListOfProduct")

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},                                                // All origins
		AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"}, // Allowing only get, just an example
	})

	// starting port 9191 server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), c.Handler(router)), nil)

}
