package main

import (
	"fmt"

	"github.com/kevinwubert/go-simple-rabbitmq-server-client/pkg/server"
)

func main() {
	err := server.Main()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
