package main

import (
	"fmt"

	"github.com/kevinwubert/go-simple-rabbitmq-server-client/pkg/producer"
)

func main() {
	err := producer.Main()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
