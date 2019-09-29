package main

import (
	"database/sql"
	"flag"
	"github.com/roman-mazur/chat-channels-example/server/db"
	"log"
)

var httpPortNumber = flag.Int("p", 8080, "HTTP port number")

func NewDbConnection() (*sql.DB, error) {
	return (&db.Connection{
		DbName:     "chat-example",
		User:       "class-admin",
		Password:   "architecture-2019",
		Host:       "localhost",
		DisableSSL: true,
	}).Open()
}

func main() {
	// Parse command line arguments. Port number may be defined with "-p" flag.
	flag.Parse()

	log.Println("Starting chat server...")

	// Create the server.
	if server, err := ComposeApiServer(HttpPortNumber(*httpPortNumber)); err == nil {
		// Start it.
		server.Start()
	} else {
		log.Fatalf("Cannot initialize chat server: %s", err)
	}
}
