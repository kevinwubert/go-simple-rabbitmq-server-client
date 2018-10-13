package producer

import (
	"github.com/kevinwubert/go-simple-rabbitmq-server-client/pkg/rabbitmq"
)

func Main() error {
	rabbitMqClient, err := rabbitmq.New()
	if err != nil {
		return err
	}

	rabbitMqClient.Publish("Hello Message")
	return nil
}
