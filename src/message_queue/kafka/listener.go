package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/mitchellh/mapstructure"
	
	"contracts"
	"helper/kafka"
	"message_queue"
)

type kafkaEventListener struct {
	consumer   sarama.Consumer
	partitions []int32
//	mapper     message_queue.EventMapper
}

func XXXNewKafkaEventListenerFromEnvironment() (message_queue.EventListener, error) {
	brokers := []string{"localhost:9092"}
	partitions := []int32{}

	if brokerList := os.Getenv("KAFKA_BROKERS"); brokerList != "" {
		brokers = strings.Split(brokerList, ",")
	}

	if partitionList := os.Getenv("KAFKA_PARTITIONS"); partitionList != "" {
		partitionStrings := strings.Split(partitionList, ",")
		partitions = make([]int32, len(partitionStrings))

		for i := range partitionStrings {
			partition, err := strconv.Atoi(partitionStrings[i])
			if err != nil {
				return nil, err
			}
			partitions[i] = int32(partition)
		}
	}

	client := <-kafka.RetryConnect(brokers, 5*time.Second)

	return NewKafkaEventListener(client, partitions)
}

func NewKafkaEventListener(client sarama.Client, partitions []int32) (message_queue.EventListener, error) {
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return nil, err
	}

	listener := &kafkaEventListener{
		consumer:   consumer,
		partitions: partitions,
//		mapper:     message_queue.NewEventMapper(),
	}

	return listener, nil
}

func (k *kafkaEventListener) Listen(events ...string) (<-chan message_queue.Event, <-chan error, error) {
	var err error

	topic := "event.created"
	results := make(chan message_queue.Event)
	errors := make(chan error)

	partitions := k.partitions
	if len(partitions) == 0 {
		partitions, err = k.consumer.Partitions(topic)
		if err != nil {
			return nil, nil, err
		}
	}

	log.Printf("topic %s has partitions: %v", topic, partitions)

	for _, partition := range partitions {
		log.Printf("consuming partition %s:%d", topic, partition)

		pConsumer, err := k.consumer.ConsumePartition(topic, partition, 0)
		if err != nil {
			log.Printf("ConsumePartition:: err= %v", err)
			return nil, nil, err
		}

		go func() {
			for msg := range pConsumer.Messages() {
				log.Printf("received message %v", msg)

				body := messageEnvelope{}
				err := json.Unmarshal(msg.Value, &body)
				if err != nil {
					errors <- fmt.Errorf("could not JSON-decode message: %v", err)
					continue
				}



				/*event, err := k.mapper.MapEvent(body.EventName, body.Payload)
				if err != nil {
					errors <- fmt.Errorf("could not map message: %v", err)
					continue
				}*/
				var event message_queue.Event 
				switch body.EventName { 
				case "event.created": 
					event = &contracts.EventCreatedEvent{}
				case "location.created": 
					event = &contracts.LocationCreatedEvent{} 
				default: 
					errors <- fmt.Errorf("unknown event type: %s", body.EventName) 
					continue 
				} 

				cfg := mapstructure.DecoderConfig{ 
					Result: event, 
					TagName: "json", 
				}

				dec, err := mapstructure.NewDecoder(&cfg)
				if err != nil {
					errors <- fmt.Errorf("could not initialize decoder for event %s: %s", body.EventName, err)
					continue 
				}

				log.Printf("body.Payload %v", body.Payload)

			    err = dec.Decode(body.Payload)
				if err != nil { 
					errors <- fmt.Errorf("could not map event %s: %s", body.EventName, err)
					continue 
				}

				results <- event
			}
		}()

		go func() {
			for err := range pConsumer.Errors() {
				errors <- err
			}
		}()
	}

	return results, errors, nil
}
/*
func (l *kafkaEventListener) Mapper() message_queue.EventMapper {
	return l.mapper
}
*/