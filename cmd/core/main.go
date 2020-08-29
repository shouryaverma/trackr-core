package main

import (
	"fmt"
	"log"

	"github.com/amaraliou/trackr-core/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error getting env %v\n", err)
	}

	server, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	server.Logger.Infof("Starting server")
	server.ListenAndServe()
}
