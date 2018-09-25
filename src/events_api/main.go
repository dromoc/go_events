package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/streadway/amqp"
	"github.com/Shopify/sarama"

    "events_api/rest"
    "configuration"
    "persistence/dblayer"
    "message_queue"
    message_queue_amqp "message_queue/amqp"
    "message_queue/kafka"
)

func main() {
	var eventEmitter message_queue.EventEmitter

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

		eventEmitter, err = message_queue_amqp.NewAMQPEventEmitter(conn, "events")
		if err != nil {
			panic(err)
		}
	case "kafka":
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
		if err != nil {
			panic(err)
		}

		eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
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
