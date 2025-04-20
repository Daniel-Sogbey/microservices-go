package main

import "log"

type Config struct {
}

const webPort = "80"

func main() {

	app := Config{}

	log.Println("Starting mail service on port", webPort)

}
