package kafka

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"helper/kafka"
	"message_queue"
)

type kafkaEventEmitter struct {
	producer sarama.SyncProducer
}

type messageEnvelope struct {
	EventName string      `json:"eventName"`
	Payload   interface{} `json:"payload"`
}

func NewKafkaEventEmitterFromEnvironment() (message_queue.EventEmitter, error) {
	brokers := []string{"localhost:9092"}

	if brokerList := os.Getenv("KAFKA_BROKERS"); brokerList != "" {
		brokers = strings.Split(brokerList, ",")
	}

	client := <-kafka.RetryConnect(brokers, 5*time.Second)
	return NewKafkaEventEmitter(client)
}

func NewKafkaEventEmitter(client sarama.Client) (message_queue.EventEmitter, error) {
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}

	emitter := kafkaEventEmitter{
		producer: producer,
	}

	return &emitter, nil
}

func (k *kafkaEventEmitter) Emit(event message_queue.Event) error {
	jsonBody, err := json.Marshal(messageEnvelope{
		event.EventName(),
		event,
	})
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "event.created",
		Value: sarama.ByteEncoder(jsonBody),
	}

	log.Printf("published message with topic %s: %v", event.EventName(), jsonBody)
	_, _, err = k.producer.SendMessage(msg)

	return err
}