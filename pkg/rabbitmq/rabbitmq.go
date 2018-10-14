package rabbitmq

import (
	"github.com/streadway/amqp"
)

type client struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

type Client interface {
	Publish(msg string) error
	Consume() (<-chan amqp.Delivery, error)
	Close()
}

func New(url string, queueName string) (Client, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return &client{
		conn: conn,
		ch:   ch,
		q:    q,
	}, nil
}

func (c *client) Publish(msg string) error {
	err := c.ch.Publish(
		"",       // exchange
		c.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	return err
}

func (c *client) Consume() (<-chan amqp.Delivery, error) {
	msgs, err := c.ch.Consume(
		c.q.Name, // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	return msgs, err
}

func (c *client) Close() {
	c.ch.Close()
	c.conn.Close()
}
