package consumer

import (
	"github.com/kevinwubert/go-simple-rabbitmq-server-client/pkg/rabbitmq"
)

func Main() error {
	rabbitMqClient, err := rabbitmq.New()
	if err != nil {
		return err
	}

	rabbitMqClient.ConsumeMessage()
	return nil
}
