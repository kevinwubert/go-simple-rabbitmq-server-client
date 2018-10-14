package config

type Config struct {
	RabbitMQURL     string
	QueueName       string
	WhitelistedCmds []string
}

func GetConfig() (Config, error) {
	return Config{
		RabbitMQURL:     "amqp://guest:guest@localhost:5672/",
		QueueName:       "task",
		WhitelistedCmds: []string{"echo"},
	}, nil
}
