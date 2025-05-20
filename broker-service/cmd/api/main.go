package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	conn *amqp.Connection
}

func main() {

	conn, err := connectToRabbitmq()
	if err != nil {
		panic(err)
	}

	app := Config{
		conn: conn,
	}

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Fatal(srv.ListenAndServe())

}

func connectToRabbitmq() (*amqp.Connection, error) {
	connection, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	if err != nil {
		log.Println("failed to connect to rabbitmq")
		return nil, err
	}

	defer connection.Close()

	return connection, err
}
