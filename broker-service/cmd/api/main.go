package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const port = "80"

type Config struct {
	Rabbit *amqp.Connection
}

func main() {
	rabbitConn, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitConn.Close()
	app := Config{
		Rabbit: rabbitConn,
	}
	log.Printf("Starting broker service on port %s\n", port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
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
