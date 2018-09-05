package main

import (
	"flag"
	"fmt"
	"log"
    "events_api/rest"
    "configuration"
    "persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", "./../configuration/config.json", "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	//RESTful API start
    fmt.Println("Start Web server: ", config.RestfulEndpoint)
	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPint, dbhandler)
	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}
}
