package main

import (
	"fmt"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/raphaelmb/go-ms-listener-service/event"
)

func main() {
	rabbitConn, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitConn.Close()

	log.Println("Listening for and consuming RabbitMQ messages...")

	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		log.Fatal(err)
	}

	if err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"}); err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	svc := "rabbitmq"
	connectionStr := fmt.Sprintf("amqp://guest:guest@%s", svc)
	counts := 0
	backOff := 1 * time.Second

	for {
		c, err := amqp.Dial(connectionStr)
		if err != nil {
			fmt.Println("RabbitMQ not ready yet.")
			counts++
		} else {
			return c, nil
		}
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off...")
		time.Sleep(backOff)
		continue
	}
}
