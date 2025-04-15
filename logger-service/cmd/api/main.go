package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"log-service/data"
	"net/http"
	"time"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://localhost:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	//connect to mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Fatal(err)
	}

	client = mongoClient

	//create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//close connection
	defer func(client *mongo.Client) {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}(client)

	app := Config{
		Models: data.New(client),
	}

	//start web server
	app.serve()
}

func (app *Config) serve() {
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.route(),
	}

	log.Println("Starting logger service on port 80")
	log.Fatal(srv.ListenAndServe())
}

func connectToMongo() (*mongo.Client, error) {
	//create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	//connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	return c, nil
}
