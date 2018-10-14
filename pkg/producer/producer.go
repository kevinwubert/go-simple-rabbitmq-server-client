package producer

import (
	"github.com/kevinwubert/go-simple-rabbitmq-server-client/pkg/config"
	"github.com/kevinwubert/go-simple-rabbitmq-server-client/pkg/rabbitmq"
	"github.com/pkg/errors"
)

func Main() error {
	config, err := config.GetConfig()
	if err != nil {
		return errors.Wrap(err, "could not get config")
	}
	rabbitMqClient, err := rabbitmq.New(config.RabbitMQURL, config.QueueName)
	if err != nil {
		return errors.Wrap(err, "could not create new rabbitmq client")
	}

	err = rabbitMqClient.Publish("Hello Message")
	if err != nil {
		return errors.Wrap(err, "could not publish message")
	}

	return nil
}
