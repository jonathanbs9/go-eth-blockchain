package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	Handlers "github.com/jonathanbs9/go-eth-blockchain/handler"
)

func main() {
	// Create a client instance to connect to our provider
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		fmt.Println(err)
	}

	// Create a mux router
	r := mux.NewRouter()

	// Define a single endpoint
	//r.Handle("/api/v1/eth/{module}", Handlers.ClientHandler{client})
	r.Handle("/api/v1/eth/{module}", Handlers.ClientHandler{client})

	log.Fatal(http.ListenAndServe(":8080", r))
}
