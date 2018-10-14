# go-simple-rabbitmq-server-client
Simple Golang template for a producer and consumer with rabbitmq following [this tutorial](https://www.rabbitmq.com/tutorials/tutorial-one-go.html), tracking job statuses in redis, and loading configuration with viper.

## Requirements
* Golang
* `make`

## Getting Started
Build the binaries with
```shell
make build
```

Install RabbitMQ
```shell
make rabbitmq
```

## Run
Run the producer
```shell
./bin/producer
```
Run the consumer
```shell
./bin/consumer
```