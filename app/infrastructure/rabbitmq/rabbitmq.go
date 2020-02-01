package rabbitmq

import (
	"log"
	"os"
	"rabbitmq-http-publisher/app/infrastructure/env"
	"github.com/streadway/amqp"
)

type Confirmable interface {
	Confirm(chan amqp.Confirmation)
}


func GetAmpqConnection() string {
	return env.GetEnvOrDefault("RabbitMq__Connection", "amqp://guest:guest@rabbitmq:5672/")
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

func DeclareExchange(ch *amqp.Channel, exchaneName string) *amqp.Channel {
	err := ch.ExchangeDeclare(
		exchaneName, // name
		"topic",     // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	env.FailOnError(err, "Error when trying declare exchange")
	return ch
}


func PublishMessage(ch *amqp.Channel, exchangeName, routeKey string, message *[]byte) error {
	c := DeclareExchange(ch, exchangeName)
	return c.Publish(exchangeName, routeKey, false, false, amqp.Publishing{ ContentType: "application/json", Body: *message});
}

