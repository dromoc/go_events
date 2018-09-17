package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/streadway/amqp"

    "events_api/rest"
    "configuration"
    "persistence/dblayer"
    "message_queue/amqp"
)

func main() {
	confPath := flag.String("conf", "./../configuration/config.json", "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	switch config.MessageBrokerType {
	case "amqp":
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		if err != nil {
			panic(err)
		}

		eventEmitter, err = msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
		if err != nil {
			panic(err)
		}

	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	//RESTful API start
    fmt.Println("Start Web server: ", config.RestfulEndpoint)
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler, eventEmitter))
}
