package rabbitmq

import (
	"github.com/offsec1/goscan/core/cli"
	"github.com/offsec1/goscan/core/utils"
	"github.com/streadway/amqp"
	"strings"
)

func Receive() {
	conn, err := amqp.Dial(utils.Config.RabbitMQUri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"GoScanCommands", // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"be01", // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			s := []string{"Received a message:", string(d.Body)}
			utils.Config.Log.LogInfo(strings.Join(s, " "))
			cli.Executor(string(d.Body))
		}
	}()

	utils.Config.Log.LogInfo("Waiting for messages. To exit press CTRL+C")
	<-forever
}
