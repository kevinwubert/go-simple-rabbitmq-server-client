package consumer

import (
	"log"

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

	msgs, err := rabbitMqClient.Consume()
	if err != nil {
		rabbitMqClient.Close()
		return errors.Wrap(err, "could not consume messages")
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL + C\n")
	<-forever

	return nil
}
