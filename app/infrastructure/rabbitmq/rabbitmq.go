package rabbitmq

import (
	"log"
	"os"
	"rossmannpl-backend-kafka-producer/app/infrastructure/env"
	"github.com/streadway/amqp"
)

func GetAmpqConnection() string {
	return env.GetEnvOrDefault("RabbitMq__Connection", "amqp://user:dCbCI41Gk5@127.0.0.1:5672/")
}

func ConnectToRabbitMq(ampqConnectionString string) *amqp.Connection {
	conn, err := amqp.Dial(ampqConnectionString)
	env.FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	env.FailOnError(err, "Failed to connect to RabbitMQ")
	return ch
}

func CloseRabbit(conn *amqp.Connection, ch *amqp.Channel) {
	err := ch.Close()
	env.FailOnError(err, "Can't close rabbit channel :( ")
	err = conn.Close()
	env.FailOnError(err, "Can't close rabbit connection :( ")

	if err != nil {
		log.Fatalf("error: %v\n", err)
		os.Exit(1)
	}
}

func DeclareQueue(ch *amqp.Channel, exchaneName, queueSufix string) <-chan amqp.Delivery {
	err := ch.ExchangeDeclare(
		exchaneName, // name
		"topic",     // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	
	env.FailOnError(err, "Failed to declare an exchange")
	q, err := ch.QueueDeclare(
		exchaneName+"-"+queueSufix, // name
		true,                       // durable
		false,                      // delete when usused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // arguments
	)
	env.FailOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,      // queue name
		"#",         // routing key
		exchaneName, // exchange
		false,
		nil,
	)
	env.FailOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name,                     // queue
		exchaneName+"-"+queueSufix, // consumer
		true,                       // auto-ack
		false,                      // exclusive
		false,                      // no-local
		false,                      // no-wait
		nil,                        // args
	)
	env.FailOnError(err, "Failed to register a consumer")
	return msgs
}
