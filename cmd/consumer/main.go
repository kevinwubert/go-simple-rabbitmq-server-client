package main

import (
	"fmt"

	"github.com/kevinwubert/go-simple-rabbitmq-server-client/pkg/consumer"
)

func main() {
	err := consumer.Main()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
