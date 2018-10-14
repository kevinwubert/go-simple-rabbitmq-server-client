package config

type Config struct {
	RabbitMQURL string
	QueueName   string
}

func GetConfig() (Config, error) {
	return Config{
		RabbitMQURL: "amqp://guest:guest@localhost:5672/",
		QueueName:   "hello",
	}, nil
}
