package main

import (
   	"fmt"
	"flag"

	"github.com/streadway/amqp"
	"github.com/Shopify/sarama"

	"booking_api/listener"
	"booking_api/rest"
	"configuration"
	"message_queue"
	message_queue_amqp "message_queue/amqp"
	"message_queue/kafka"
	"persistence/dblayer"
)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var eventListener message_queue.EventListener
	var eventEmitter message_queue.EventEmitter

	confPath := flag.String("conf", "./../configuration/config_booking.json", "flag to set the path to the configuration json file")
	flag.Parse()

	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Printf("MessageBrokerType: %s", config.MessageBrokerType)

	switch config.MessageBrokerType {
	case "amqp":
		conn, err := amqp.Dial(config.AMQPMessageBroker)
		panicIfErr(err)

		eventListener, err = message_queue_amqp.NewAMQPEventListener(conn, "events", "booking")
		panicIfErr(err)

		eventEmitter, err = message_queue_amqp.NewAMQPEventEmitter(conn, "events")
		panicIfErr(err)
	case "kafka":
		conf := sarama.NewConfig()
		conf.Producer.Return.Successes = true
		conn, err := sarama.NewClient(config.KafkaMessageBrokers, conf)
		panicIfErr(err)

		eventListener, err = kafka.NewKafkaEventListener(conn, []int32{})
		panicIfErr(err)

		eventEmitter, err = kafka.NewKafkaEventEmitter(conn)
		panicIfErr(err)
	default:
		panic("Bad message broker type: " + config.MessageBrokerType)
	}

    fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	processor := listener.EventProcessor{eventListener, dbhandler}
	go processor.ProcessEvents()

	fmt.Println("Start Web server: ", config.RestfulEndpoint)
	rest.ServeAPI(config.RestfulEndpoint, dbhandler, eventEmitter)
}